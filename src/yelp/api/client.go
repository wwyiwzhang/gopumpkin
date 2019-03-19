package yelper

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type HTTPClient interface {
	Get(url string, kwargs map[string]string) (HTTPResponse, error)
	Post(url string, kwargs map[string]string, payload string) error
}

type yelpHTTPClient struct {
	token  string
	client *http.Client
}

var token string

func NewYelpHTTPClient() HTTPClient {
	return &yelpHTTPClient{
		token:  token,
		client: &http.Client{},
	}
}

type HTTPResponse interface {
	Decode(Businesses) (Businesses, error)
}

// yelpHTTPResponse implements HTTPResponse
type yelpHTTPResponse struct {
	body   io.ReadCloser
	status int
}

func (yresp *yelpHTTPResponse) Decode(v Businesses) (Businesses, error) {
	b, err := ioutil.ReadAll(yresp.body)
	if err != nil {
		return Businesses{}, fmt.Errorf("failed to read response: %v", err)
	}
	err = json.Unmarshal(b, &v)
	if err != nil {
		return Businesses{}, fmt.Errorf("failed to parse JSON encoded data: %v", err)
	}
	return v, nil
}

func (y *yelpHTTPClient) Get(url string, kwargs map[string]string) (HTTPResponse, error) {
	query := []string{}
	for k, v := range kwargs {
		query = append(query, strings.Join([]string{k, v}, "="))
	}
	fullquery := strings.Join(query, "&")
	fullURL := fmt.Sprintf("%s?%s", url, fullquery)

	req, err := http.NewRequest("GET", fullURL, nil)
	req.Header.Add("Authorization", "Bearer "+y.token)

	resp, err := y.client.Do(req)
	if err != nil {
		return &yelpHTTPResponse{
			body: nil, status: resp.StatusCode,
		}, fmt.Errorf("failed to make GET request at %s: %v", fullURL, err)
	}
	return &yelpHTTPResponse{body: resp.Body, status: resp.StatusCode}, nil
}

func (y *yelpHTTPClient) Post(url string, kwargs map[string]string, payload string) error {
	return fmt.Errorf("Not implemented")
}

func init() {
	token = os.Getenv("DEVELOPER_API_TOKEN")
}
