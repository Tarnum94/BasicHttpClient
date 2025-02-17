package gohttp

import (
	"net/http"
	"testing"
)

func TestSetRequestHeaders(t *testing.T) {
	//Initialization
	client := httpClient{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/barbie")
	commonHeaders.Set("User-Agent", "cool-http-client")
	client.defaultHeaders = commonHeaders

	//Execution
	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-Id", "ABC-123")
	requestHeaders.Set("Content-Type", "application/xml")
	requestHeaders.Add("Content-Type", "application/json")
	finalHeaders := client.setRequestHeaders(requestHeaders)

	//Validation
	if len(finalHeaders) != 3 {
		t.Error("we expected 3 headers")
	}

	if len(finalHeaders.Values("Content-Type")) != 2 {
		t.Error("wrong number of header values")
	}

	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Error("invalid cintent type received")
	}

	if finalHeaders.Get("User-Agent") != "cool-http-client" {
		t.Error("invalid user-id received")
	}

	if finalHeaders.Get("X-Request-Id") != "ABC-123" {
		t.Error("adding custom headers failed")
	}
}
