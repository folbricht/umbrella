# Umbrella API client

[![GoDoc](https://godoc.org/github.com/folbricht/umbrella?status.svg)](https://godoc.org/github.com/folbricht/umbrella)

Umbrella is a client library as well as a command-line tool for the [APIs](https://docs.umbrella.com/developer) provided by Cisco Umbrella, formerly OpenDNS. It supports all endpoints available in the [Investigate API](https://docs.umbrella.com/developer/investigate-api/) and the [Enforcement API](https://docs.umbrella.com/developer/enforcement-api/). Contributions are welcome.

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

### Tools

The command-line tools mainly exists for testing purposes and to show how the library can be used. There is one sub-command for each API endpoint and the two tools currently available can be installed with:

```shell
go get -u "github.com/folbricht/umbrella/cmd/investigate"
go get -u "github.com/folbricht/umbrella/cmd/enforcement"
```

Accessing the Umbrella Investigate API requires an API token which can be passed into the tool either by environment variable `UMBRELLA_KEY` or by command-line option `-key`. For the Enforcement API, a customer key is required which can be provided by environment variable `CUSTOMER_KEY` or the `-key` option. The tools support multiple sub-commands, each of which represents a specific API call. Examples of how to use the tools:

```shell
UMBRELLA_KEY=... investigate domain-timeline ihaveabadreputation.com
```

```shell
investigate -key <KEY> domain-categorization -showlabels ihaveabadreputation.com
```

```shell
investigate -key <KEY> domain-history A ihaveabadreputation.com
```

```shell
enforcement -key <KEY> list-all-domains
```

#### `investigate` Subcommands

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

#### `enforcement` Subcommands

- `list-domains` - List domains currently on the blocklist (includes pagination)
- `list-all-domains` - List all domains currently on the blocklist
- `delete-domain` - Remove a domain from the blocklist

See `enforcement <command> -h` for details on any command and available options.
