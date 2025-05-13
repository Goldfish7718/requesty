package types

type Environment struct {
	EnvironmentName string    `json:"environmentName"`
	BaseUrl         string    `json:"baseUrl"`
	Requests        []Request `json:"requests"`
}

type Request struct {
	ReqType string `json:"type"`
	Route   string `json:"route"`
}
