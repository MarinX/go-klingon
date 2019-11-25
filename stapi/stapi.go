package stapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

const (
	// UserAgent is HTTP User-Agent Header
	UserAgent = "gostapi/1.0"
	// BaseURL for STAPI
	BaseURL = "http://stapi.co/api"
	// BaseVersion for STAPI
	BaseVersion = "v1/rest"
)

// Client manages communication with the STAPI Rest API.
type Client struct {
	client *http.Client
	apiKey string

	Character CharacterService
}

// New returns new STAPI client
func New(apiKey string, client *http.Client) *Client {
	if client == nil {
		client = http.DefaultClient
	}
	cl := &Client{
		client: client,
		apiKey: apiKey,
	}

	return cl
}

// Get performs a GET request for the given path and saves the result in the
// given resource.
func (c *Client) Get(path string, options interface{}, resource interface{}) error {
	return c.execute(http.MethodGet, path, options, resource)
}

// Post performs a POST request for the given path and saves the result in the
// given resource.
func (c *Client) Post(path string, options interface{}, resource interface{}) error {
	return c.execute(http.MethodPost, path, options, resource)
}

// newRequest creates a new http request by given method, url and query params
func (c *Client) newRequest(method, urlPath string, options interface{}) (*http.Request, error) {

	apiURL, err := url.Parse(urlPath)
	if err != nil {
		return nil, err
	}

	urlParams := url.Values{}
	urlParams.Set("apiKey", c.apiKey)

	if options != nil {
		queryParams, err := query.Values(options)
		if err != nil {
			return nil, err
		}

		for k, v := range queryParams {
			for _, values := range v {
				urlParams.Add(k, values)
			}
		}
	}

	var body io.Reader
	if method != http.MethodGet {
		body = strings.NewReader(urlParams.Encode())
	} else {
		apiURL.RawQuery = urlParams.Encode()
	}

	req, err := http.NewRequest(method, apiURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", UserAgent)

	if method != http.MethodGet {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	return req, nil
}

// execute performs a web request to STAPI with the given method (GET,POST) and relative path.
func (c *Client) execute(method, resourcePath string, options interface{}, resource interface{}) error {
	req, err := c.newRequest(method, resourcePath, options)
	if err != nil {
		return err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error http status api %d", resp.StatusCode)
	}

	if resource != nil {
		err := json.NewDecoder(resp.Body).Decode(resource)
		if err != nil {
			return err
		}
	}
	return nil
}

// BuildURL builds the relative path
func BuildURL(path string) string {
	return fmt.Sprintf("%s/%s/%s", BaseURL, BaseVersion, path)
}
