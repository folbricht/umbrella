package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/folbricht/umbrella"
)

const asUsage = `as <IP>

Find AS information for an IP.`

func as(key string, args []string) error {
	flags := flag.NewFlagSet("as", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, asUsage)
		flags.PrintDefaults()
	}
	flags.Parse(args)
	if flags.NArg() < 1 {
		return notEnoughArgs
	}
	if flags.NArg() > 1 {
		return tooManyArgs
	}
	ip := flags.Arg(0)
	client := umbrella.NewInvestigate(key)
	results, err := client.ASForIP(ip)
	if err != nil {
		return err
	}
	return printJSON(results)
}

const prefixesUsage = `prefixes <ASN>

Find CIDR and Geo information for an ASN.`

func prefixes(key string, args []string) error {
	flags := flag.NewFlagSet("prefixes", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, prefixesUsage)
		flags.PrintDefaults()
	}
	flags.Parse(args)
	if flags.NArg() < 1 {
		return notEnoughArgs
	}
	if flags.NArg() > 1 {
		return tooManyArgs
	}
	asn, err := strconv.Atoi(flags.Arg(0))
	if err != nil {
		return err
	}
	client := umbrella.NewInvestigate(key)
	results, err := client.PrefixesForASN(asn)
	if err != nil {
		return err
	}
	return printJSON(results)
}
