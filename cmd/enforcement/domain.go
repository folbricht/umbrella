package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/folbricht/umbrella"
)

const listDomainsUsage = `list-domains [options]

Show blocklist items.`

func getDomains(key string, args []string) error {
	var (
		page, limit int
	)
	flags := flag.NewFlagSet("list-domains", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, listDomainsUsage)
		flags.PrintDefaults()
	}
	flags.IntVar(&page, "page", 1, "Requested result page")
	flags.IntVar(&limit, "limit", 200, "Limit the number of results")
	flags.Parse(args)
	if flags.NArg() > 0 {
		return errTooManyArgs
	}
	client := umbrella.NewEnforcement(key)
	opts := umbrella.QueryOptions{
		"page":  {strconv.Itoa(page)},
		"limit": {strconv.Itoa(limit)},
	}
	domains, err := client.GetDomains(opts)
	if err != nil {
		return err
	}
	return printJSON(domains)
}

const listAllDomainsUsage = `list-all-domains

Show complete list of blocklist items.`

func getAllDomains(key string, args []string) error {
	flags := flag.NewFlagSet("list-all-domains", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, listAllDomainsUsage)
		flags.PrintDefaults()
	}
	flags.Parse(args)
	if flags.NArg() > 0 {
		return errTooManyArgs
	}
	client := umbrella.NewEnforcement(key)
	domains, err := client.GetAllDomains()
	if err != nil {
		return err
	}
	return printJSON(domains)
}

const deleteDomainUsage = `delete-domain <DOMAIN-OR-ID>

Remove a domain from the blocklist.`

func deleteDomain(key string, args []string) error {
	flags := flag.NewFlagSet("delete-domain", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, deleteDomainUsage)
		flags.PrintDefaults()
	}
	flags.Parse(args)
	if flags.NArg() < 1 {
		return errNotEnoughArgs
	}
	if flags.NArg() > 1 {
		return errTooManyArgs
	}
	domain := flags.Arg(0)
	client := umbrella.NewEnforcement(key)
	return client.DeleteDomain(domain)
}
