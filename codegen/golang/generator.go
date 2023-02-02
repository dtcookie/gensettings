package golang

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"

	"github.com/dtcookie/dynatrace/gensettings/codegen"
	"github.com/dtcookie/dynatrace/gensettings/codegen/camel"
	"github.com/dtcookie/dynatrace/gensettings/reflection"

	"golang.org/x/tools/imports"
)

type generator struct {
	library  *reflection.Library
	schemaID string
}

func NewGenerator(library *reflection.Library, schemaID string) codegen.Generator {
	return &generator{library: library, schemaID: schemaID}
}

func (me *generator) Write(folder string) error {
	if err := os.WriteFile(folder+"/schema.json", me.library.JSON, 0644); err != nil {
		panic(err)
	}
	for name := range me.library.Types {
		t := me.library.Types[name]
		switch t.Kind {
		case reflection.EnumKind:
			props := reflection.Properties{}
			for _, property := range t.Properties {
				property.Name = PropertyName(property.Name)
				props[property.Name] = property
			}
			t.Properties = props
			t.Name = EnumTypeName(t.ID)
		case reflection.StructKind:
			t.Name = StructTypeName(t.ID)
		}
	}

	for _, t := range me.library.AllTypes() {
		if t.Kind == reflection.ListKind {
			if t.Elem.Kind == reflection.StructKind {
				replacement := &reflection.Type{
					ID:   Plural(t.Elem.Unref().ID),
					Name: Plural(t.Elem.Unref().Name),
					Kind: reflection.SliceAliasKind,
					Elem: t,
				}
				me.library.Rewire(t, replacement)
				me.library.Define(replacement, true)
			}
		} else if t.Kind == reflection.SetKind {
			if t.Elem.Kind == reflection.StructKind {
				replacement := &reflection.Type{
					ID:   Plural(t.Elem.Unref().ID),
					Name: Plural(t.Elem.Unref().Name),
					Kind: reflection.SetAliasKind,
					Elem: t,
				}
				me.library.Rewire(t, replacement)
				me.library.Define(replacement, true)
			}
		}
	}

	for _, t := range me.library.AllTypes() {
		if t.Kind == reflection.StructKind {
			for propertyName, property := range t.Properties {
				if property.Type.Kind == reflection.SliceAliasKind || property.Type.Kind == reflection.SetAliasKind {
					property.Type.Properties = reflection.Properties{propertyName: nil}
				}
			}
		}
	}

	files := map[string][]CodeContributor{}
	for _, t := range me.library.Types {
		contributors, err := me.TranslateType(folder, t)
		if err != nil {
			return err
		}
		for _, contributor := range contributors {
			fileName := contributor.FileName()
			if storedContributors, exists := files[fileName]; exists {
				files[fileName] = append(storedContributors, contributor)
			} else {
				files[fileName] = []CodeContributor{contributor}
			}
		}
	}
	for fileName, contributors := range files {
		buf := new(bytes.Buffer)
		goFile := &GoFile{
			Package: packageName(me.schemaID),
			Name:    fileName,
			Imports: []string{},
		}
		sort.Slice(contributors, func(i, j int) bool {
			kindI := contributors[i].Kind()
			kindJ := contributors[j].Kind()
			if kindI < kindJ {
				return true
			} else if kindI > kindJ {
				return false
			}
			return strings.Compare(contributors[i].SortName(), contributors[j].SortName()) < 0
		})
		for _, contributor := range contributors {
			data, err := contributor.Bytes()
			if err != nil {
				return err
			}
			if imports := contributor.Imports(); len(imports) > 0 {
				goFile.Imports = append(goFile.Imports, imports...)
			}
			if _, err := buf.Write(data); err != nil {
				return err
			}
		}
		goFile.Code = buf.String()
		data, err := postProcess(goFile)
		if err != nil {
			return err
		}
		os.MkdirAll(path.Join(folder, "settings"), os.ModePerm)
		if err := os.WriteFile(path.Join(path.Join(folder, "settings"), fileName), data, 0644); err != nil {
			return err
		}
	}
	return nil
}

func (me *generator) TranslateType(folder string, t *reflection.Type) ([]CodeContributor, error) {
	switch t.Kind {
	case reflection.EnumKind:
		return []CodeContributor{NewEnum(t)}, nil
	case reflection.StructKind:
		return []CodeContributor{NewStruct(t)}, nil
	case reflection.SliceAliasKind, reflection.SetAliasKind:
		item := ""
		for k := range t.Properties {
			item = k
		}
		var singular = ""
		var err error
		if singular, err = Singular(item); err != nil {
			singular = camel.Strip(TypeName(t.Elem.Unref()))
		}
		if strings.HasSuffix(singular, "_item") {
			singular = "item"
		}
		if t.Kind == reflection.SetAliasKind {
			return []CodeContributor{NewSetAlias(t, singular)}, nil
		}
		return []CodeContributor{NewSliceAlias(t, singular)}, nil
	default:
		return nil, fmt.Errorf("unsupported kind %v", t.Kind)
	}
}

func postProcess(goFile *GoFile) ([]byte, error) {
	tmpl, err := template.New("GO_FILE").Parse(GO_FILE_TEMPLATE)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, goFile)
	if err != nil {
		return nil, err
	}
	data := buf.Bytes()

	processedData, err := imports.Process(goFile.Name, data, nil)
	if err != nil {
		for lineNum, line := range strings.Split(string(data), "\n") {
			fmt.Println(lineNum+1, ":", line)
		}

	}
	return processedData, err
}
