package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/dtcookie/dynatrace/gensettings/codegen/golang"
	"github.com/dtcookie/dynatrace/gensettings/collections"
	"github.com/dtcookie/dynatrace/gensettings/reflection"
	"github.com/dtcookie/dynatrace/gensettings/schema"
)

type Credentials struct {
	EnvironmentURL string
	APIToken       string
}

var CREDENTIALS = Credentials{}

func LoadCredentials() {
	CREDENTIALS.EnvironmentURL = os.Getenv("DYNATRACE_ENV_URL")
	CREDENTIALS.APIToken = os.Getenv("DYNATRACE_API_TOKEN")
	if CREDENTIALS.EnvironmentURL == "" {
		fmt.Println("Environment variable DYNATRACE_ENV_URL not specified. Example: https://#######.live.dynatrace.com")
		os.Exit(0)
	}
	if CREDENTIALS.APIToken == "" {
		fmt.Println("Environment variable DYNATRACE_API_TOKEN not specified. That token requires permission settings.Read")
		os.Exit(0)
	}
}

const SchemaFolder = "schemas"
const SettingsFolder = "settings"

func packageFolder(schemaID string) string {
	schemaID = strings.ReplaceAll(schemaID, ":", ".")
	schemaID = strings.ReplaceAll(schemaID, "-", "")
	return path.Join(strings.Split(schemaID, ".")...)
}

func GET[T any](url string, target T) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Api-Token "+CREDENTIALS.APIToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, target)
}

func GETX[T any](url string, target T) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Api-Token "+CREDENTIALS.APIToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(data), json.Unmarshal(data, target)
}

func MarshalTo(v any, path string) error {
	data, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return err
	}
	m := map[string]any{}
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	data, err = json.MarshalIndent(m, "", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

type Schema struct {
	Definition schema.Definition
	JSON       []byte
}

func checkScope(scopes collections.Set[string], elem string) bool {
	if len(scopes) > 2 {
		return false
	}
	if len(scopes) == 1 {
		return scopes.Contains(elem)
	}
	if len(scopes) == 2 {
		return scopes.Contains("environment") && scopes.Contains(elem)
	}
	return false
}

var scopeIds = map[string]string{
	"SERVICE":            "serviceId",
	"PROCESS_GROUP":      "processGroupId",
	"HOST_GROUP":         "hostGroupId",
	"HOST":               "hostId",
	"MOBILE_APPLICATION": "applicationId",
	"APPLICATION":        "applicationId",
	"DISK":               "diskId",
}

func main() {
	LoadCredentials()
	var schemaList schema.List
	if err := GET(CREDENTIALS.EnvironmentURL+"/api/v2/settings/schemas", &schemaList); err != nil {
		panic(err)
	}
	schemata := map[string]*Schema{}
	for _, schemaStub := range schemaList.Items {
		// if schemaStub.SchemaID != "builtin:cloud.cloudfoundry" {
		// 	continue
		// }
		// if !strings.HasPrefix(schemaStub.SchemaID, "builtin:") {
		// 	fmt.Println(schemaStub.SchemaID)
		// }

		var schemaDefinition schema.Definition
		data, err := GETX(CREDENTIALS.EnvironmentURL+"/api/v2/settings/schemas/"+url.PathEscape(schemaStub.SchemaID), &schemaDefinition)
		if err != nil {
			panic(err)
		}
		m := map[string]any{}
		if err := json.Unmarshal([]byte(data), &m); err != nil {
			panic(err)
		}
		data2, err := json.MarshalIndent(m, "", "\t")
		if err != nil {
			panic(err)
		}
		schemata[schemaDefinition.SchemaID] = &Schema{schemaDefinition, data2}
	}

	if err := os.RemoveAll(SettingsFolder); err != nil {
		panic(err)
	}

	if err := os.MkdirAll(SettingsFolder, os.ModeDir); err != nil {
		panic(err)
	}

	for schemaID, schema := range schemata {
		fmt.Println(schemaID)
		resolver := NewTypeResolver(&schema.Definition, schema.JSON)
		if len(schema.Definition.Enums) > 0 {
			for name := range schema.Definition.Enums {
				resolver.Resolve("#/enums/" + name)
			}
		}
		if len(schema.Definition.Types) > 0 {
			for name := range schema.Definition.Types {
				resolver.Resolve("#/types/" + name)
			}
		}
		typeName := "Settings"
		rootType := reflection.Type{
			ID:         typeName,
			Kind:       reflection.StructKind,
			Properties: reflection.Properties{},
		}
		if len(schema.Definition.AllowedScopes) > 1 || !schema.Definition.AllowedScopes.Contains("environment") {
			comment := "The scope of this setting (" + schema.Definition.AllowedScopes.ToString() + ")"
			if schema.Definition.AllowedScopes.Contains("environment") {
				comment = comment + ". Omit this property if you want to cover the whole environment."
			}
			scopeProperty := &reflection.Property{
				Name:     "scope",
				Type:     &reflection.TypeString,
				Comment:  comment,
				Optional: schema.Definition.AllowedScopes.Contains("environment"),
				Scope:    "scope",
			}
			for k, v := range scopeIds {
				if checkScope(schema.Definition.AllowedScopes, k) {
					scopeProperty.Name = v
					scopeProperty.Scope = v
					scopeProperty.Comment = "The scope of this settings. If the settings should cover the whole environment, just don't specify any scope."
					break
				}
			}
			rootType.Properties[scopeProperty.Name] = scopeProperty
		}
		resolver.Library.Define(&rootType, false)
		for propertyName, propertyDefinition := range schema.Definition.Properties {
			comment := propertyDefinition.Description
			if len(comment) == 0 {
				comment = propertyDefinition.Documentation
			}
			if len(comment) == 0 {
				comment = propertyDefinition.DisplayName
			}
			property := &reflection.Property{
				Name:            propertyName,
				Type:            resolver.resolvePropertyType(propertyName, propertyDefinition),
				Comment:         comment,
				Optional:        propertyDefinition.Nullable || propertyDefinition.Precondition != nil || (propertyDefinition.MinObjects != nil && *propertyDefinition.MinObjects == 0),
				OptionalComment: OptionalComment(propertyDefinition),
				Precondition:    translatePrecondition(propertyDefinition),
				Constraints:     translateConstraints(propertyDefinition),
				Sensitive:       propertyDefinition.Type == "secret",
			}
			if propertyDefinition.Nullable {
				property.Type = property.Type.Pointer()
			}
			rootType.Properties[propertyName] = property
		}

		curPackageFolder := path.Join(SettingsFolder, packageFolder(schemaID))

		if err := os.MkdirAll(curPackageFolder, os.ModeDir); err != nil {
			panic(err)
		}

		generator := golang.NewGenerator(resolver.Library, schemaID)
		if err := generator.Write(curPackageFolder); err != nil {
			panic(err)
		}
	}

	hide(hide)
}

func hide(v interface{}) {}
