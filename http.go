package umbrella

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// HTTPClient is an interface that allows the use of different clients to
// execute HTTP requests, for example to add logging or for testing.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func do(client HTTPClient, req *http.Request, out interface{}) error {
	if client == nil {
		client = http.DefaultClient
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	defer ioutil.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return NewUnexpectedResponse(resp.StatusCode, resp.Body)
	}
	if out != nil {
		err = json.NewDecoder(resp.Body).Decode(out)
	}
	return err
}
