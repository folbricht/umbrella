Umbrella API client
===================

Umbrella is a client library as well as a command-line tool for the [APIs](https://docs.umbrella.com/developer) provided by Cisco Umbrella, formerly OpenDNS. It supports all endpoints available in the [Investigate API](https://docs.umbrella.com/developer/investigate-api/) at this point. The [Enforcement API](https://docs.umbrella.com/developer/enforcement-api/) will be implemented at a later time. Contributions are welcome.

### Library

Import the library.
```go
import "github.com/folbricht/umbrella"
```

Create a client for the Investigate API. Requires an API access token.
```go
client := umbrella.NewInvestigate(key)
```

Query available categories.
```go
categories, err := client.DomainCategories()
if err != nil {
  return err
}
for key, value := range categories {
  // do something
}
```

Query categorization of multiple domains.
```go
domains := []string{"umbrella.com", "ihaveabadreputation.com"}
categorizations, err := client.DomainCategorizations(domains...)
if err != nil {
  return err
}
for domain, categorization := range categorizations {
  // do something
}
```

### Tool

The command-line tool mainly exists for testing purposes and to show how the library can be used. There is one sub-command for each API endpoint and it can be installed with:

```
go get -u "github.com/folbricht/umbrella"
```

Accessing the Umbrella API requires an API token which can be passed into the tool either by environment variable `UMBRELLA_KEY` or by command-line option `-key`. The tool supports multiple sub-commands, each of which represents a specific API call. Examples of how to use the tool:

```
UMBRELLA_KEY=... investigate domain-timeline ihaveabadreputation.com
```

```
investigate -key <KEY> domain-categorization -showlabels ihaveabadreputation.com
```

```
investigate -key <KEY> domain-history A ihaveabadreputation.com
```

#### Subcommands

- `domain-categories` - List category IDs and Labels
- `domain-categorization` - Categorization for a single domain
- `domain-categorizations` - Categorization of multiple domains
- `domain-timeline` - Show the timeline of a domain
- `domain-volume` - Query volume of a domain
- `search` - Perform a pattern search
- `co-occurrences` - Find domains that were queried around the same time by the same client
- `related` - Find domains related to a domain
- `security` - Show available security information for a domain
- `domain-history` - Query the history of a domain+type
- `ip-history` - Query the history of a ip+type
- `as` - Query the Autonomous System information for an IP
- `prefixes` - Query CIDR and Geo information for an ASN
- `whois-email` - Query the domains registered for a single email
- `whois-emails` - Query the domains registered for multiple emails
- `latest-malicious` - Query the (malicious) domains associated with an IP
- `top-million` - Show the top most popular domains (up to 1 million)
- `samples` - List samples associated with an IP, domain, or URL
- `sample` - Show information about a single sample by file hash
- `sample-artifacts` - Show information about artifacts associated with a sample
- `sample-connections` - Show information about connections associated with a sample
- `sample-behaviors` - List indicators associated with a sample

See `investigate <command> -h` for details on any command and available options.

### Links

- GoDoc documentation for the library - https://godoc.org/github.com/folbricht/umbrella
