package client

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// Client is an interface which provides method signatures for a HTTP client
type Client interface {
	Get(url string) ([]byte, int, error)
	Post(url string, data io.Reader, headers map[string]string) ([]byte, int, error)
}

type standardClient struct {
	client *http.Client
}

var (
	_ Client = (*standardClient)(nil)
)

// NewClient returns an instanced HTTP client
func NewClient() (Client, error) {
	return &standardClient{
		client: &http.Client{
			Timeout: time.Second * 30,
		},
	}, nil
}

// makeRequest will make any HTTP request and also sends common data required for each request
func (c *standardClient) makeRequest(method string, url string, body io.Reader, headers map[string]string) ([]byte, int, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, -100, err
	}

	request.Header.Set("Content-Type", "application/json")
	for header := range headers {
		request.Header.Set(header, headers[header])
	}
	response, err := c.client.Do(request)
	if err != nil {
		return nil, -101, err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, -102, err
	}

	return responseBody, response.StatusCode, nil
}

// Get makes a GET request to a given URL
func (c *standardClient) Get(url string) ([]byte, int, error) {
	return c.makeRequest("GET", url, nil, nil)
}

// Post makes a POST request to a givn URL
func (c *standardClient) Post(url string, data io.Reader, headers map[string]string) ([]byte, int, error) {
	return c.makeRequest("POST", url, data, headers)
}
