package umbrella

import (
	"encoding/json"
	"fmt"
	"math"
	"net/url"
)

// RR is a resource record used in history query results.
type RR struct {
	Name  string `json:"name"`
	TTL   int    `json:"ttl"`
	Type  string `json:"type"`
	RR    string `json:"rr"`
	Class string `json:"class"`
}

// DomainHistoryRR is a record of a time period in which resource records were valid.
type DomainHistoryRR struct {
	FirstSeen string `json:"first_seen"`
	LastSeen  string `json:"last_seen"`
	RRs       []RR   `json:"rrs"`
}

// DomainHistoryFeatures is used in history queries and contains historical
// information as well as features of a domain.
type DomainHistoryFeatures struct {
	Age           int      `json:"age"`
	TTLsMin       int      `json:"ttls_min"`
	TTLsMax       int      `json:"ttls_max"`
	TTLsMean      float64  `json:"ttls_mean"`
	TTLsMedian    float64  `json:"ttls_median"`
	TTLsStdDev    float64  `json:"ttls_stddev"`
	CountryCodes  []string `json:"country_codes"`
	CountryCount  int      `json:"country_count"`
	ASNs          []int    `json:"asns"`
	ASNsCount     int      `json:"asns_count"`
	Prefixes      []string `json:"prefixes"`
	PrefixesCount float64  `json:"prefixes_count"`
	RIPs          int      `json:"rips"`
	DivRIPs       float64  `json:"div_rips"` // can be a float or string "NaN"
	Locations     []struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"locations"`
	LocationsCount  int     `json:"locations_count"`
	GeoDistanceSum  float64 `json:"geo_distance_sum"`
	GeoDistanceMean float64 `json:"geo_distance_mean"`
	NonRoutable     bool    `json:"non_routable"`
	MailExchanger   bool    `json:"mail_exchanger"`
	CName           bool    `json:"cname"`
	FFCandidate     bool    `json:"ff_candidate"`
	RIPsStability   float64 `json:"rips_stability"`
	BaseDomain      string  `json:"base_domain"`
	IsSubdomain     bool    `json:"is_subdomain"`
}

// MarshalJSON is needed to handle div_rips which can be either a float or "NaN"
func (u DomainHistoryFeatures) MarshalJSON() ([]byte, error) {
	type Alias DomainHistoryFeatures
	var DivRIPs interface{} = u.DivRIPs
	if math.IsNaN(u.DivRIPs) {
		DivRIPs = "NaN"
	}
	b, err := json.Marshal(&struct {
		DivRIPs interface{} `json:"div_rips"`
		Alias
	}{
		DivRIPs: DivRIPs,
		Alias:   (Alias)(u),
	})
	return b, err
}

// UnmarshalJSON is needed to handle div_rips which can be either a float or "NaN"
func (u *DomainHistoryFeatures) UnmarshalJSON(data []byte) error {
	type Alias DomainHistoryFeatures
	aux := struct {
		DivRIPs interface{} `json:"div_rips"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if v, ok := aux.DivRIPs.(float64); ok {
		u.DivRIPs = v
	} else {
		u.DivRIPs = math.NaN()
	}
	return nil
}

// DomainHistory of resource records for a domain by record type
type DomainHistory struct {
	RRsTF    []DomainHistoryRR     `json:"rrs_tf"`
	Features DomainHistoryFeatures `json:"features"`
}

// IPHistory contains historical RRs as well as features of an IP
type IPHistory struct {
	RRs      []RR `json:"rrs"`
	Features struct {
		RRCount   int     `json:"rr_count"`
		LD2Count  int     `json:"ld2_count"`
		LD3Count  int     `json:"ld3_count"`
		LD21Count int     `json:"ld2_1_count"`
		LD22Count int     `json:"ld2_2_count"`
		DivLD2    float64 `json:"div_ld2"`
		DivLD3    float64 `json:"div_ld3"`
		DivLD21   float64 `json:"div_ld2_1"`
		DivLD22   float64 `json:"div_ld2_2"`
	} `json:"features"`
}

// GetDomainHistory returns the history of a domain for a given record type, typ
// must be one of A, NS, MX, TXT or CNAME.
func (c Investigate) GetDomainHistory(typ, domain string) (DomainHistory, error) {
	u := &url.URL{
		Path: fmt.Sprintf("dnsdb/name/%s/%s.json", typ, domain),
	}
	var out DomainHistory
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out, err
}

// GetIPHistory returns the history of an IP. Type must be one of "A" or "NS".
func (c Investigate) GetIPHistory(typ, ip string) (IPHistory, error) {
	u := &url.URL{
		Path: fmt.Sprintf("dnsdb/ip/%s/%s.json", typ, ip),
	}
	var out IPHistory
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out, err
}
