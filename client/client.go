package client

import (
	"io/ioutil"
	"log"
	"net/http"
)

type APIClient struct {
	config     *Configuration
	httpClient *http.Client
}

func NewClient(config *Configuration) *APIClient {
	return &APIClient{
		config:     config,
		httpClient: &http.Client{Timeout: config.Timeout},
	}
}

func (apiClient *APIClient) Call(apiReq *APIRequest) (string, error) {

	req, err := buildHttpReq(apiReq, *apiClient.config)
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

func buildHttpReq(apiReq *APIRequest, config Configuration) (*http.Request, error) {

	endpoint := config.URL + apiReq.path
	req, err := http.NewRequest(apiReq.method, endpoint, nil)
	if err != nil {
		log.Println("buildHttpReq:", err)
		return nil, err
	}

	q := req.URL.Query()
	for key, value := range config.Keys {
		q.Add(key, value)
	}
	for param, value := range apiReq.params {
		q.Add(param, value)
	}

	req.URL.RawQuery = q.Encode()
	return req, nil
}
