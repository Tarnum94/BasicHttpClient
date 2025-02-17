package gohttp

import (
	"net/http"
)

type httpClient struct {
	defaultHeaders http.Header
}

func New() HttpClient {
	client := &httpClient{}
	return client
}

type HttpClient interface {
	SetHeaders(headers http.Header)

	Get(string, http.Header) (*http.Response, error)
	Post(string, http.Header, interface{}) (*http.Response, error)
	Put(string, http.Header, interface{}) (*http.Response, error)
	Patch(string, http.Header, interface{}) (*http.Response, error)
	Delete(string, http.Header) (*http.Response, error)
}

func (c *httpClient) SetHeaders(headers http.Header) {
	c.defaultHeaders = headers
}

func (c *httpClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do("GET", url, headers, nil)
}

func (c *httpClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do("POST", url, headers, body)
}

func (c *httpClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do("PUT", url, headers, body)
}

func (c *httpClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do("Patch", url, headers, body)
}

func (c *httpClient) Delete(url string, headers http.Header) (*http.Response, error) {
	return c.do("DELETE", url, headers, nil)
}
