package umbrella

import (
	"bytes"
	"encoding/json"
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

// DefaultInvestigateURL is the Umbrella service's default API URL
const DefaultInvestigateURL = "https://investigate.api.umbrella.com"

// NewInvestigate returns a new client for the Investigate API using the
// default URL.
func NewInvestigate(token string) Investigate {
	u, _ := url.Parse(DefaultInvestigateURL)
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
	return do(c.Client, req, out)
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
	req.Header.Set("Content-Type", "application/json")
	return do(c.Client, req, out)
}
