package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/folbricht/umbrella"
)

const relatedUsage = `related <domain>

Find related domains.`

func related(key string, args []string) error {
	flags := flag.NewFlagSet("related", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, relatedUsage)
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
	results, err := client.Related(domain)
	if err != nil {
		return err
	}
	return printJSON(results)
}
