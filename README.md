Umbrella API client
===================

Umbrella is a client library as well as a command-line tool for the [APIs](https://docs.umbrella.com/developer) provided by Cisco Umbrella, formerly OpenDNS. It is still incomplete with only parts of the [Investigate API](https://docs.umbrella.com/developer/investigate-api/) being implemented at this point. Contributions are welcome.

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

The command-line tool mainly exists for testing purposes and to show how the library can be used. It can be installed with:

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

See `investigate <command> -h` for details on any command and available options.
