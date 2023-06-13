package golang

type PropertyNameOverride struct {
	PropertyName string
	HCLTag       string
}

var Overrides = map[string]map[string]map[string]PropertyNameOverride{
	"builtin:problem.notifications": {
		"AnsibleTowerNotification": {
			"AcceptAnyCertificate": {
				PropertyName: "Insecure",
				HCLTag:       "insecure",
			},
		},
	},
}

func Override(schemaID string, contributor CodeContributor) {
	var s *Struct
	switch c := contributor.(type) {
	case *Struct:
		s = c
	default:
		return
	}
	if overridesForSchema, ok := Overrides[schemaID]; ok {
		if overridesForType, ok := overridesForSchema[s.Name]; ok {
			for _, property := range s.Properties {
				if override, ok := overridesForType[property.Name]; ok {
					if len(override.PropertyName) > 0 {
						property.Name = override.PropertyName
					}
					if len(override.HCLTag) > 0 {
						property.HCLTag = override.HCLTag
					}
				}
			}
		}
	}
}
