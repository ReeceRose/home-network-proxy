package client

import (
	"crypto/tls"
	"crypto/x509"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ReeceRose/home-network-proxy/internal/consts"
	"github.com/ReeceRose/home-network-proxy/internal/utils"
)

// Client is an interface which provides method signatures for a HTTP client
type Client interface {
	Get(url string, useCustomCert bool) ([]byte, int, error)
	Post(url string, data io.Reader, useCustomCert bool) ([]byte, int, error)
}

type standardClient struct {
	client           *http.Client
	customCertClient *http.Client
}

var (
	_ Client = (*standardClient)(nil)
)

// NewClient returns an instanced HTTP client
func NewClient() (Client, error) {
	certDir := utils.GetVariable(consts.CERT_DIR)
	caCert, err := ioutil.ReadFile(certDir + "/" + utils.GetVariable(consts.CLIENT_CERT))
	if err != nil {
		panic(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	return &standardClient{
		client: &http.Client{
			Timeout: time.Second * 30,
		},
		customCertClient: &http.Client{
			Timeout: time.Second * 30,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: caCertPool,
				},
			},
		},
	}, nil
}

// makeRequest will make any HTTP request and also sends common data required for each request
func (c *standardClient) makeRequest(method string, url string, body io.Reader, useCustomCert bool) ([]byte, int, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, -100, err
	}

	request.Header.Set("Content-Type", "application/json")
	var response *http.Response
	if useCustomCert {
		response, err = c.customCertClient.Do(request)
	} else {
		response, err = c.client.Do(request)
	}
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
func (c *standardClient) Get(url string, useCustomCert bool) ([]byte, int, error) {
	return c.makeRequest("GET", url, nil, useCustomCert)
}

// Post makes a POST request to a givn URL
func (c *standardClient) Post(url string, data io.Reader, useCustomCert bool) ([]byte, int, error) {
	return c.makeRequest("POST", url, data, useCustomCert)
}
