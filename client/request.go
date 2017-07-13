package client

// APIRequest contains the necessary information for a http request execution
type APIRequest struct {
	method string
	path   string
	params map[string]string
}

// BaseAPIReq builds a basic APIRequest
func BaseAPIReq() *APIRequest {
	apiReq := &APIRequest{method: "GET", params: make(map[string]string)}
	return apiReq
}

// WithPath sets the path to be used in the http request execution
func (apiReq *APIRequest) WithPath(path string) *APIRequest {
	apiReq.path = "/" + path
	return apiReq
}

// AddParam adds a new parameter to be used int the http request execution
func (apiReq *APIRequest) AddParam(param string, value string) *APIRequest {
	apiReq.params[param] = value
	return apiReq
}
