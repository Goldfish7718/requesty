package types

type Project struct {
	ProjectName string    `json:"projectName"`
	BaseUrl     string    `json:"baseUrl"`
	Requests    []Request `json:"requests"`
}

type Request struct {
	ReqType string `json:"type"`
	Route   string `json:"route"`
}
