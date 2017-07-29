package client

import (
	"io/ioutil"
	"log"
	"net/http"
)

// APIClient is the generic client that serves as the base for specific client implementations
type APIClient struct {
	config     *Configuration
	httpClient *http.Client
}

// NewClient builds a new APIClient from a configuration
func NewClient(config *Configuration) *APIClient {
	return &APIClient{
		config:     config,
		httpClient: &http.Client{Timeout: config.timeout},
	}
}

// Call execute http requests using APIRequest. The answer is a raw json.
func (apiClient *APIClient) Call(apiReq *APIRequest) (string, error) {

	req, err := buildHTTPReq(apiReq, *apiClient.config)
	if err != nil {
		log.Printf("call: %s", err.Error())
		return "", err
	}

	log.Println(req.URL.String())
	log.Flags()

	httpClient := apiClient.httpClient
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Printf("call: %s", err.Error())
		return "", err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("call: %s", err.Error())
		return "", err
	}

	defer resp.Body.Close()

	return string(body), nil
}

func buildHTTPReq(apiReq *APIRequest, config Configuration) (*http.Request, error) {

	endpoint := config.url + apiReq.path
	req, err := http.NewRequest(apiReq.method, endpoint, nil)
	if err != nil {
		log.Println("buildHTTPReq:", err)
		return nil, err
	}

	q := req.URL.Query()
	for key, value := range config.keys {
		q.Add(key, value)
	}
	for param, value := range apiReq.params {
		q.Add(param, value)
	}

	req.URL.RawQuery = q.Encode()
	return req, nil
}
