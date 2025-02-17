package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
)

func (c *httpClient) do(method, url string, customHeaders http.Header, body interface{}) (*http.Response, error) {
	client := http.Client{}

	fullHeaders := c.setRequestHeaders(customHeaders)
	requestBody, err := c.extractRequestBody(fullHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewReader(requestBody))

	if err != nil {
		return nil, errors.New("unable to create a new request")
	}

	request.Header = fullHeaders

	return client.Do(request)
}

func (c *httpClient) setRequestHeaders(customHeaders http.Header) http.Header {
	result := make(http.Header)

	for header, values := range customHeaders {
		if len(values) > 0 {
			for _, value := range values {
				result.Add(header, value)
			}
		}
	}

	for header, values := range c.defaultHeaders {
		if len(result.Get(header)) > 0 {
			continue
		}
		if len(values) > 0 {
			for _, value := range values {
				result.Add(header, value)
			}
		}
	}

	return result
}

func (c *httpClient) extractRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		//TODO: I am not sure if I like that I would almost prefer to panic here or return emty body
		return nil, nil
	}

	switch contentType {
	case "application/json":
		return json.Marshal(body)

	case "application/xml":
		return xml.Marshal(body)

	default:
		return json.Marshal(body)
	}
}
