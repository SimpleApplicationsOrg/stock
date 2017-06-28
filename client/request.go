package client

type APIRequest struct {
	method string
	path   string
	params map[string]string
}

func BaseAPIReq() *APIRequest {
	apiReq := &APIRequest{method: "GET", params: make(map[string]string)}
	return apiReq
}

func (apiReq *APIRequest) WithPath(path string) *APIRequest {
	apiReq.path = path
	return apiReq
}

func (apiReq *APIRequest) AddParam(param string, value string) *APIRequest {
	apiReq.params[param] = value
	return apiReq
}
