package graphql_service

type GraphqlRequest struct {
	Query     string     `json:"query"`
	Headers   []Header   `json:"headers"`
	Variables []Variable `json:"variables"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Variable struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}
