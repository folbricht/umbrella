package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/folbricht/umbrella"
)

const coOccurrencesUsage = `co-occurrences <domain>

Find co-occurrences for a domains.`

func coOccurrences(key string, args []string) error {
	flags := flag.NewFlagSet("coOccurrences", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, coOccurrencesUsage)
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
	results, err := client.CoOccurrences(domain)
	if err != nil {
		return err
	}
	return printJSON(results)
}
