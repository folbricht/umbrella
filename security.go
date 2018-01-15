package umbrella

import (
	"fmt"
	"net/url"
)

// SecurityInfo contains mutliple scores and security features.
type SecurityInfo struct {
	DGAScore               float64         `json:"dga_score"`
	Perplexity             float64         `json:"perplexity"`
	Entropy                float64         `json:"entropy"`
	SecureRank2            float64         `json:"securerank2"`
	PageRank               float64         `json:"pagerank"`
	ASNScore               float64         `json:"asn_score"`
	PrefixScore            float64         `json:"prefix_score"`
	RIPScore               float64         `json:"rip_score"`
	Popularity             float64         `json:"popularity"`
	GeoDiversity           [][]interface{} `json:"geodiversity"`
	GeoDiversityNormalized [][]interface{} `json:"geodiversity_normalized"`
	TLDGeoDiversity        [][]interface{} `json:"tld_geodiversity"`
	GeoScore               float64         `json:"geoscore"`
	KSTest                 float64         `json:"ks_test"`
	Attack                 string          `json:"attack"`
	ThreatType             string          `json:"threat_type"`
	Found                  bool            `json:"found"`
}

// Security queries available security information for a domain.
func (c Investigate) Security(domain string) (SecurityInfo, error) {
	u := &url.URL{
		Path: fmt.Sprintf("security/name/%s", domain),
	}
	var out SecurityInfo
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out, err
}
