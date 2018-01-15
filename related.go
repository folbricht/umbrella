package umbrella

import (
	"fmt"
	"net/url"
)

// RelatedDomains holds a list of related domains, ie. seen 60s before/after
// domain is queries
type RelatedDomains struct {
	TB1   []interface{} `json:"tb1"`
	Found bool          `json:"found"` // true if there are results
}

// Related returns a structure describing related domains.
func (c Investigate) Related(domain string) (RelatedDomains, error) {
	u := &url.URL{
		Path: fmt.Sprintf("links/name/%s.json", domain),
	}
	var out RelatedDomains
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out, err
}
