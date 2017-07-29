package client

// APIRequest contains the necessary information for a http request execution
type APIRequest struct {
	method string
	path   string
	params map[string]string
}

// NewAPIReq builds a basic APIRequest using method GET as default
func NewAPIReq() *APIRequest {
	apiReq := &APIRequest{method: "GET", params: make(map[string]string)}
	return apiReq
}

// Method sets the method to be used in the http request execution
func (apiReq *APIRequest) Method(method string) *APIRequest {
	apiReq.method = method
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
