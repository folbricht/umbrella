package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/folbricht/umbrella"
)

const securityUsage = `security <domain>

Query security information for a domain.`

func security(key string, args []string) error {
	flags := flag.NewFlagSet("security", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, securityUsage)
		flags.PrintDefaults()
	}
	flags.Parse(args)
	if flags.NArg() < 1 {
		return notEnoughArgs
	}
	if flags.NArg() > 1 {
		return tooManyArgs
	}
	domain := flags.Arg(0)
	client := umbrella.NewInvestigate(key)
	results, err := client.Security(domain)
	if err != nil {
		return err
	}
	return printJSON(results)
}
