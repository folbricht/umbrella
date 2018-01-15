package umbrella

import (
	"fmt"
	"net/url"
	"time"
)

// SearchResults contains a list of matches when performning a pattern search.
type SearchResults struct {
	Expression        string `json:"expression"`        // Query Regex
	TotalResults      int    `json:"totalResults"`      // Number of results
	MoreDataAvailable bool   `json:"moreDataAvailable"` // True if TotalResults > Limit
	Limit             int    `json:"limit"`             // Query result limit
	Matches           []struct {
		Name               string    `json:"name"` // Domain name
		FirstSeen          int64     `json:"firstSeen"`
		FirstSeenISO       time.Time `json:"firstSeenISO"`
		SecurityCategories []string  `json:"securityCategories"`
	} `json:"matches"`
}

// Search performs a pattern search.
func (c Investigate) Search(expression string, opts QueryOptions) (SearchResults, error) {
	u := &url.URL{
		Path:     fmt.Sprintf("search/%s", expression),
		RawQuery: url.Values(opts).Encode(),
	}
	var out SearchResults
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out, err
}
