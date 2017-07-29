package client_test

import (
	"fmt"
	"github.com/SimpleApplicationsOrg/stock/client"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

const (
	okResponse    = `{"test":"response"}`
	errorResponse = `{"fake error json string"}`
	validAPIKey   = "validApiKey"
	invalidAPIKey = "invalidApiKey"
)

func buildTestServer(okResponse string) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Query().Get("apikey") == validAPIKey {
			fmt.Fprint(w, okResponse)

		} else {
			fmt.Fprint(w, errorResponse)
		}

	}))
	return ts
}

func buildTimeoutServer(okResponse string) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		time.Sleep(100 * time.Millisecond)
		fmt.Fprint(w, okResponse)

	}))
	return ts
}

func TestCallSuccess(t *testing.T) {
	ts := buildTestServer(okResponse)
	defer ts.Close()
	config := client.NewConfiguration(ts.URL).AddKey("apikey", validAPIKey)
	api := client.NewClient(config)

	resp, err := api.Call(client.NewAPIReq())
	if err != nil {
		t.Errorf("error on call: %s", err)
	}

	if resp != okResponse {
		t.Errorf("received incorrect response: %s", resp)
	}
}

func TestCallError(t *testing.T) {
	ts := buildTestServer(okResponse)
	defer ts.Close()
	config := client.NewConfiguration(ts.URL).AddKey("apikey", invalidAPIKey)
	api := client.NewClient(config)

	resp, _ := api.Call(client.NewAPIReq())

	if resp != errorResponse {
		t.Errorf("should receive error")
	}
}

func TestCallApiWithTimeout(t *testing.T) {
	ts := buildTimeoutServer(okResponse)
	defer ts.Close()

	config := client.NewConfiguration(ts.URL).
		AddKey("apikey", validAPIKey).
		WithTimeout(10 * time.Millisecond)
	api := client.NewClient(config)

	_, err := api.Call(client.NewAPIReq())

	if err == nil || !strings.Contains(err.Error(), "Client.Timeout") {
		t.Errorf("should have timeout")
	}
}

func Example() {
	config := client.NewConfiguration("https://api.com").
		AddKey("apiKey", "value").
		WithTimeout(3 * time.Second)

	newClient := client.NewClient(config)

	apiReq := client.NewAPIReq().
		WithPath("endpointPath").
		AddParam("search", "example")

	response, err := newClient.Call(apiReq)
	if err != nil {
		fmt.Printf("error calling api: %s", err.Error())
		return
	}
	fmt.Println(response)
}
