package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/folbricht/umbrella"
)

const domainHistoryUsage = `domain-history <type> <domain>

History of a domain's resource records by type.`

func domainHistory(key string, args []string) error {
	flags := flag.NewFlagSet("domainHistory", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, domainHistoryUsage)
		flags.PrintDefaults()
	}
	flags.Parse(args)
	if flags.NArg() < 2 {
		return notEnoughArgs
	}
	if flags.NArg() > 2 {
		return tooManyArgs
	}
	typ := flags.Arg(0)
	domain := flags.Arg(1)
	client := umbrella.NewInvestigate(key)
	results, err := client.GetDomainHistory(typ, domain)
	if err != nil {
		return err
	}
	return printJSON(results)
}

const ipHistoryUsage = `ip-history <type> <ip>

History of a domain's resource records by type.`

func ipHistory(key string, args []string) error {
	flags := flag.NewFlagSet("ipHistory", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, ipHistoryUsage)
		flags.PrintDefaults()
	}
	flags.Parse(args)
	if flags.NArg() < 2 {
		return notEnoughArgs
	}
	if flags.NArg() > 2 {
		return tooManyArgs
	}
	typ := flags.Arg(0)
	ip := flags.Arg(1)
	client := umbrella.NewInvestigate(key)
	results, err := client.GetIPHistory(typ, ip)
	if err != nil {
		return err
	}
	return printJSON(results)
}
