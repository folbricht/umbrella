package umbrella

import (
	"fmt"
	"net/url"
)

// Volume contains the number of DNS queries made per hour
type Volume struct {
	Dates   []int `json:"dates"`   // Dates for which data is returned, Epoch
	Queries []int `json:"queries"` // Number of DNS queries per hour
}

// DomainVolume returns the query volume for a given domain.
func (c Investigate) DomainVolume(domain string, opts QueryOptions) (Volume, error) {
	u := &url.URL{
		Path:     fmt.Sprintf("domains/volume/%s", domain),
		RawQuery: url.Values(opts).Encode(),
	}
	var out Volume
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out, err
}
