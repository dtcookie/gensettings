package schema

type Stub struct {
	Description         string `json:"description"`
	SchemaID            string `json:"schemaId"`
	LatestSchemaVersion string `json:"latestSchemaVersion"`
	DisplayName         string `json:"displayName"`
}
