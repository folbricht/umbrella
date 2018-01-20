package umbrella

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
)

// Enforcement API Client
type Enforcement struct {
	CustomerKey string // Enforcement API token
	BaseURL     *url.URL
	Client      HTTPClient // Defaults to http.DefaultClient if nil
}

// DefaultEnforcementURL is the Umbrella service's default URL for the Enforcement API
const DefaultEnforcementURL = "https://s-platform.api.opendns.com"

// NewEnforcement returns a new client for the Enforcement API using the
// default URL.
func NewEnforcement(key string) Enforcement {
	u, _ := url.Parse(DefaultEnforcementURL)
	return Enforcement{
		CustomerKey: key,
		BaseURL:     u,
	}
}

// Get is a convenience function that adds authentication to a request and
// parses the returned JSON into out (should be a pointer).
func (c Enforcement) Get(location string, out interface{}) error {
	u, err := url.Parse(location)
	if err != nil {
		return err
	}
	q := u.Query()
	q.Set("customerKey", c.CustomerKey)
	u.RawQuery = q.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return err
	}
	return do(c.Client, req, out)
}

// Post is a convenience function that adds authentication to a request, POSTs
// the content of 'in' as JSON and parses the returned JSON into out. Both, in
// and out should be pointers.
func (c Enforcement) Post(location string, in, out interface{}) error {
	u, err := url.Parse(location)
	if err != nil {
		return err
	}
	q := u.Query()
	q.Set("customerKey", c.CustomerKey)
	b := new(bytes.Buffer)
	if err = json.NewEncoder(b).Encode(in); err != nil {
		return err
	}
	req, err := http.NewRequest("POST", u.String(), b)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	return do(c.Client, req, out)
}

// Delete is a convenience function that adds authentication to a request and
// makes a DELETE call to the API.
func (c Enforcement) Delete(location string) error {
	u, err := url.Parse(location)
	if err != nil {
		return err
	}
	q := u.Query()
	q.Set("customerKey", c.CustomerKey)
	u.RawQuery = q.Encode()
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return err
	}
	return do(c.Client, req, nil)
}
