package umbrella

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// DomainBlocklist represents a blocklist as returned by the Enforcement API.
// Contains metadata used for pagination.
type DomainBlocklist struct {
	Meta Meta                  `json:"meta"`
	Data []DomainBlocklistItem `json:"data"`
}

// Meta is used for pagination in the Enforcement API. Unfortunately, 'prev' and
// 'next' require special handling since those are either of type bool (false)
// or a string with the URL.
type Meta struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}

// UnmarshalJSON is needed to handle prev/next as they come back as bool or string
func (u *Meta) UnmarshalJSON(data []byte) error {
	type Alias Meta
	aux := struct {
		Prev interface{} `json:"prev"`
		Next interface{} `json:"next"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if v, ok := aux.Prev.(string); ok {
		u.Prev = v
	}
	if v, ok := aux.Next.(string); ok {
		u.Next = v
	}
	return nil
}

// DomainBlocklistItem represents a domain on a blocklist as returned by the
// Enforcement API
type DomainBlocklistItem struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	LastSeenAt time.Time `json:"lastSeenAt"`
}

// UnmarshalJSON is needed to convert the Epoch time into proper time.Time
func (u *DomainBlocklistItem) UnmarshalJSON(data []byte) error {
	type Alias DomainBlocklistItem
	aux := struct {
		LastSeenAt int64 `json:"lastSeenAt"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	u.LastSeenAt = time.Unix(aux.LastSeenAt, 0)
	return nil
}

// GetDomains returns the results of a single blocklist query, including
// pagination. Supports 'page' and 'limit' options.
func (c Enforcement) GetDomains(opts QueryOptions) (DomainBlocklist, error) {
	u := &url.URL{
		Path:     "1.0/domains",
		RawQuery: url.Values(opts).Encode(),
	}
	var out DomainBlocklist
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out, err
}

// GetAllDomains returns the complete list of items on the blocklist. Will
// do multiple queries with pagination if necessary.
func (c Enforcement) GetAllDomains() ([]DomainBlocklistItem, error) {
	var items []DomainBlocklistItem
	u := c.BaseURL.ResolveReference(&url.URL{Path: "1.0/domains"}).String()

	// Keep making queries and appending to the item list unti next == "" (last page)
	for u != "" {
		var bl DomainBlocklist
		if err := c.Get(u, &bl); err != nil {
			return items, err
		}
		u = bl.Meta.Next
		items = append(items, bl.Data...)
	}
	return items, nil
}

// DeleteDomain removes a domain from a blocklist. The 'domain' parameter can
// be either the ID (as resturned by the API) or the domain name.
func (c Enforcement) DeleteDomain(domain string) error {
	u := &url.URL{
		Path: fmt.Sprintf("1.0/domains/%s", domain),
	}
	return c.Delete(c.BaseURL.ResolveReference(u).String())
}
