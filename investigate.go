package umbrella

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

// QueryOptions are passed in URL query parameters
type QueryOptions map[string][]string

// Investigate API Client
type Investigate struct {
	APIToken string // Umbrella API token
	BaseURL  *url.URL
	Client   HTTPClient // Defaults to http.DefaultClient if nil
}

// DefaultURL is the Umbrella service's default API URL
const DefaultURL = "https://investigate.api.umbrella.com"

// HTTPClient is an interface that allows the use of different clients to
// execute HTTP requests, for example to add logging or for testing.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// NewInvestigate returns a new client for the Investigate API using the
// default URL.
func NewInvestigate(token string) Investigate {
	u, _ := url.Parse(DefaultURL)
	return Investigate{
		APIToken: token,
		BaseURL:  u,
	}
}

// Get is a convenience function that adds authentication to a request and
// parses the returned JSON into out (should be a pointer).
func (c Investigate) Get(url string, out interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	return c.do(req, out)
}

// Post is a convenience function that adds authentication to a request, POSTs
// the content of 'in' as JSON and parses the returned JSON into out. Both, in
// and out should be pointers.
func (c Investigate) Post(url string, in, out interface{}) error {
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(in); err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+c.APIToken)
	return c.do(req, out)
}

func (c Investigate) do(req *http.Request, out interface{}) error {
	var client HTTPClient = http.DefaultClient
	if c.Client != nil {
		client = c.Client
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	defer ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return NewUnexpectedResponse(resp.StatusCode, resp.Body)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}
