package umbrella

import (
	"fmt"
	"net/url"
)

// AS (Autonomous System) information
type AS struct {
	CreationDate string `json:"creation_date"`
	IR           int    `json:"ir"`
	Description  string `json:"description"`
	ASN          int    `json:"asn"`
	CIDR         string `json:"cidr"`
}

// Prefix contains information about CIDR blocks associated with an IP and their
// location.
type Prefix struct {
	CIDR string `json:"cidr"`
	Geo  struct {
		CountryName string `json:"country_name"`
		CountryCode string `json:"country_code"`
	} `json:"geo"`
}

// ASForIP returns a list of Autonomous Systems that an IP address is associated
// with.
func (c Investigate) ASForIP(ip string) ([]AS, error) {
	u := &url.URL{
		Path: fmt.Sprintf("bgp_routes/ip/%s/as_for_ip.json", ip),
	}
	var out []AS
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out, err
}

// PrefixesForASN returns CIDR and Geo information linked to an ASN.
func (c Investigate) PrefixesForASN(asn int) ([]Prefix, error) {
	u := &url.URL{
		Path: fmt.Sprintf("bgp_routes/asn/%d/prefixes_for_asn.json", asn),
	}
	var out []Prefix
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out, err
}
