package umbrella

import (
	"net/url"
)

// TopMillion returns the top most popular domains (up to 1 million). Typically
// the 'limit' option is used to restrict the number of results and limit
// memory consumption.
func (c Investigate) TopMillion(opts QueryOptions) ([]string, error) {
	u := &url.URL{
		Path:     "topmillion",
		RawQuery: url.Values(opts).Encode(),
	}
	var out []string
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out, err
}
