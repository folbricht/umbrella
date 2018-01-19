package umbrella

import (
	"fmt"
	"net/url"
)

// Behavior describes a single indicator associated with a sample.
type Behavior struct {
	Name       string   `json:"name"`
	Title      string   `json:"title"`
	Hits       int      `json:"hits"`
	Confidence int      `json:"confidence"`
	Severity   int      `json:"severity"`
	Tags       []string `json:"tags"`
	Threat     int      `json:"threat"`
	Category   []string `json:"category"`
}

// Artifact represents a single artifact, typically associated with a sample
type Artifact struct {
	SHA256    string `json:"sha256"`
	SHA1      string `json:"sha1"`
	MD5       string `json:"md5"`
	MagicType string `json:"magicType"`
	Size      int    `json:"size"`
	FirstSeen int    `json:"firstSeen"`
	LastSeen  int    `json:"lastSeen"`
	Visible   bool   `json:"visible"`
	AVResults []struct {
		Signature string `json:"signature"`
		Product   string `json:"product"`
	} `json:"avresults"`
	Behaviors []Behavior `json:"behaviors"`
}

// Sample represents a single sample, returned by either /sample or /samples.
type Sample struct {
	Artifact
	ThreatScore int `json:"threatScore"`
}

// Connection contains information about a single network connection made
// by a sample.
type Connection struct {
	Name               string   `json:"name"`
	FirstSeen          int      `json:"firstSeen"`
	LastSeen           int      `json:"lastSeen"`
	SecurityCategories []string `json:"securityCategories"`
	Attacks            []string `json:"attacks"`
	ThreatTypes        []string `json:"threatTypes"`
	Type               string   `json:"type"`
	IPs                []string `json:"ips"`
	URLs               []string `json:"urls"`
}

// SampleFull hold all information about a single sample, including its artifacts
// and network connections.
type SampleFull struct {
	Sample
	Artifacts struct {
		Pagignation
		Artifacts []Artifact `json:"artifacts"`
	} `json:"artifacts"`
	Samples struct {
		Pagignation
		Samples []Sample `json:"samples"`
	} `json:"samples"`
	Connections struct {
		Pagignation
		Connections []Connection `json:"connections"`
	} `json:"connections"`
}

// SampleList holds a list of samples matching a query by domain, URL, or IP
type SampleList struct {
	Pagignation
	Query   string   `json:"query"`
	Samples []Sample `json:"samples"`
}

// ArtifactList is returned when querying artifacts related to a sample, includes
// pagignation.
type ArtifactList struct {
	Pagignation
	Artifacts []Artifact `json:"artifacts"`
}

// ConnectionList holds a list of network connections make by a samples, including
// pagination.
type ConnectionList struct {
	Pagignation
	Connections []Connection `json:"connections"`
}

// Pagignation is used in several query responses to show if there are
// additional records available as well as show the limit and current position.
type Pagignation struct {
	TotalResults      int  `json:"totalResults"`
	MoreDataAvailable bool `json:"moreDataAvailable"`
	Limit             int  `json:"limit"`
	Offset            int  `json:"offset"`
}

// GetSamples returns all samples associated with an IP, a domain or a URL.
// Available options are 'limit', 'offset', and 'sortby'.
func (c Investigate) GetSamples(term string, opts QueryOptions) (SampleList, error) {
	fmt.Printf("samples/%s", url.QueryEscape(term))
	u := &url.URL{
		Path:     fmt.Sprintf("samples/%s", term),
		RawPath:  fmt.Sprintf("samples/%s", url.QueryEscape(term)),
		RawQuery: url.Values(opts).Encode(),
	}
	var out SampleList
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out, err
}

// GetSample returns complete information about a single sample, identified by its hash.
// Hash can be either a SHA256, SHA1, or MD5. Supports 'limit' and 'offset' options.
func (c Investigate) GetSample(hash string, opts QueryOptions) (SampleFull, error) {
	u := &url.URL{
		Path:     fmt.Sprintf("sample/%s", hash),
		RawQuery: url.Values(opts).Encode(),
	}
	var out SampleFull
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out, err
}

// GetSampleArtifacts returns artifacts associated with a sample, identified by
// its hash. Hash can be either a SHA256, SHA1, or MD5. Supports 'limit' and
// 'offset' options.
func (c Investigate) GetSampleArtifacts(hash string, opts QueryOptions) (ArtifactList, error) {
	u := &url.URL{
		Path:     fmt.Sprintf("sample/%s/artifacts", hash),
		RawQuery: url.Values(opts).Encode(),
	}
	var out ArtifactList
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out, err
}

// GetSampleConnections returns connections associated with a sample, identified by
// its hash. Hash can be either a SHA256, SHA1, or MD5. Supports 'limit' and
// 'offset' options.
func (c Investigate) GetSampleConnections(hash string, opts QueryOptions) (ConnectionList, error) {
	u := &url.URL{
		Path:     fmt.Sprintf("sample/%s/connections", hash),
		RawQuery: url.Values(opts).Encode(),
	}
	var out ConnectionList
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out, err
}

// GetSampleBehaviors returns indicators associated with a sample, identified by
// its hash. Hash can be either a SHA256, SHA1, or MD5.
func (c Investigate) GetSampleBehaviors(hash string) ([]Behavior, error) {
	u := &url.URL{
		Path: fmt.Sprintf("sample/%s/behaviors", hash),
	}
	var out []Behavior
	err := c.Get(c.BaseURL.ResolveReference(u).String(), &out)
	return out, err
}
