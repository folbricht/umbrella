package umbrella

import (
	"fmt"
	"net/url"
)

// CoOccurrenceResults contains a list of domains that were accessed at the same time.
type CoOccurrenceResults struct {
	PFS2  []interface{} `json:"pfs2"`
	Found bool          `json:"found"` // true if there are results
}

// CoOccurrences returns a structure describing other domains that were queried
// around the same time as domain.
func (c Investigate) CoOccurrences(domain string) (CoOccurrenceResults, error) {
	u := &url.URL{
		Path: fmt.Sprintf("recommendations/name/%s.json", domain),
	}
	var out CoOccurrenceResults
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out, err
}
