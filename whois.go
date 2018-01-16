package umbrella

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// WHOIS information for an email address.
type WHOIS struct {
	TotalResults      int  `json:"totalResults"`
	MoreDataAvailable bool `json:"moreDataAvailable"`
	Limit             int  `json:"limit"`
	Domains           []struct {
		Domain  string `json:"domain"`
		Current bool   `json:"current"`
	} `json:"domains"`
}

// WHOISEmail returns the domains for a single registrant.
func (c Investigate) WHOISEmail(limit int, email string) (WHOIS, error) {
	u := &url.URL{
		Path:     fmt.Sprintf("whois/emails/%s", email),
		RawQuery: url.Values{"limit": {strconv.Itoa(limit)}}.Encode(),
	}
	var out map[string]WHOIS
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out[email], err
}

// WHOISEmails returns the domains for multiple registrants in a map.
func (c Investigate) WHOISEmails(limit int, emails ...string) (map[string]WHOIS, error) {
	u := &url.URL{
		Path: "whois/emails",
		RawQuery: url.Values{
			"emailList": {strings.Join(emails, ",")},
			"limit":     {strconv.Itoa(limit)},
		}.Encode(),
	}
	var out map[string]WHOIS
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out, err
}
