package umbrella

import (
	"fmt"
	"net/url"
)

// Categorization of a domain. If the query included the "showlabels", the
// categories will contain full string values describing them. If "showlabels"
// wasn't used, those will category numbers (as strings).
type Categorization struct {
	Status             int      `json:"status"`
	SecurityCategories []string `json:"security_categories"`
	ContentCategories  []string `json:"content_categories"`
}

// DomainCategorization returns the categorization for a single domain. If
// showlabels is true, the response will contain the string labels instead
// of category numbers.
func (c Investigate) DomainCategorization(domain string, showLabels bool) (Categorization, error) {
	u := &url.URL{
		Path: fmt.Sprintf("domains/categorization/%s", domain),
	}
	if showLabels {
		u.RawQuery = url.Values{"showLabels": {""}}.Encode()
	}
	var out map[string]Categorization
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out[domain], err
}

// DomainCategorizations returns the categorizations of multiple domains in a map.
func (c Investigate) DomainCategorizations(domains ...string) (map[string]Categorization, error) {
	u := &url.URL{
		Path: "domains/categorization",
	}
	var out map[string]Categorization
	err := c.Post(c.BaseURL.ResolveReference(u).String(), &domains, &out)
	return out, err
}

// DomainCategories returns a mapping of category numbers to their string labels.
func (c Investigate) DomainCategories() (map[int]string, error) {
	u := &url.URL{
		Path: "domains/categories",
	}
	var out map[int]string
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out, err
}
