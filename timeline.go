package umbrella

import (
	"fmt"
	"net/url"
)

// TimelineEvent represents a change to category, threat type, etc at a specific
// point in time.
type TimelineEvent struct {
	Categories  []string `json:"categories"`
	Attacks     []string `json:"attacks"`
	ThreatTypes []string `json:"threatTypes"`
	Timestamp   int64    `json:"timestamp"` // Epoch
}

// Timeline contains the history of security events associated with a domain.
type Timeline []TimelineEvent

// DomainTimeline returns the timeline for a domain.
func (c Investigate) DomainTimeline(domain string) (Timeline, error) {
	u := &url.URL{
		Path: fmt.Sprintf("timeline/%s", domain),
	}
	var out Timeline
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out, err
}
