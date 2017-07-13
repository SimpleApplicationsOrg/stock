package client

import (
	"fmt"
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
	config := NewConfiguration(ts.URL).AddKey("apikey", validAPIKey)
	api := NewClient(config)

	resp, err := api.Call(BaseAPIReq())
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
	config := NewConfiguration(ts.URL).AddKey("apikey", invalidAPIKey)
	api := NewClient(config)

	resp, _ := api.Call(BaseAPIReq())

	if resp != errorResponse {
		t.Errorf("should receive error")
	}

}

func TestCallApiWithTimeout(t *testing.T) {
	ts := buildTimeoutServer(okResponse)
	defer ts.Close()

	config := NewConfiguration(ts.URL).
		AddKey("apikey", validAPIKey).
		WithTimeout(10 * time.Millisecond)
	api := NewClient(config)

	_, err := api.Call(BaseAPIReq())

	if err == nil || !strings.Contains(err.Error(), "Client.Timeout") {
		t.Errorf("should have timeout")
	}

}
