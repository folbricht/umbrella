package umbrella

import (
	"net/url"
	"time"
)

// Event is a malware event that can be POST'ed to the Enforcement API.
type Event struct {
	AlertTime       time.Time `json:"alertTime"`
	DeviceID        string    `json:"deviceId"`
	DeviceVersion   string    `json:"deviceVersion"`
	DstDomain       string    `json:"dstDomain"`
	DstURL          string    `json:"dstUrl"`
	EventTime       time.Time `json:"eventTime"`
	ProtocolVersion string    `json:"protocolVersion"`
	ProviderName    string    `json:"providerName"`
}

// PostEvent sends a single event to the Enforcement API.
func (c Enforcement) PostEvent(e Event) error {
	u := &url.URL{
		Path: "1.0/events",
	}
	return c.Post(c.BaseURL.ResolveReference(u).String(), &e, nil)
}

// PostEvents sends multiple events to the Enforcement API.
func (c Enforcement) PostEvents(e ...Event) error {
	u := &url.URL{
		Path: "1.0/events",
	}
	return c.Post(c.BaseURL.ResolveReference(u).String(), &e, nil)
}
