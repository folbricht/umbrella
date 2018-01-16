package umbrella

import (
	"fmt"
	"net/url"
)

// Domain is returned when querying malicious domains for an IP.
type Domain struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// LatestMalicious returns a list of malicious domains associated with an IP.
func (c Investigate) LatestMalicious(ip string) ([]Domain, error) {
	u := &url.URL{
		Path: fmt.Sprintf("ips/%s/latest_domains", ip),
	}
	var out []Domain
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out, err
}
