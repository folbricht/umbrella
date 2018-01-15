package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/folbricht/umbrella"
)

const searchUsage = `search [options] <pattern>

Perform a pattern search on the API`

func search(key string, args []string) error {
	var (
		start           string
		includeCategory string
	)
	flags := flag.NewFlagSet("search", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, searchUsage)
		flags.PrintDefaults()
	}
	flags.StringVar(&start, "start", "-1days", "Start time")
	flags.StringVar(&includeCategory, "includeCategory", "false", "Include security categories")
	flags.Parse(args)
	if flags.NArg() < 1 {
		return notEnoughArgs
	}
	if flags.NArg() > 1 {
		return tooManyArgs
	}
	expression := flags.Arg(0)
	client := umbrella.NewInvestigate(key)
	opts := umbrella.QueryOptions{
		"start":           {start},
		"includeCategory": {includeCategory},
	}
	results, err := client.Search(expression, opts)
	if err != nil {
		return err
	}
	return printJSON(results)
}
