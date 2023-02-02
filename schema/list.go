package schema

type List struct {
	Description string `json:"description"`
	TotalCount  int64  `json:"totalCount"`
	Items       []Stub `json:"items"`
}
