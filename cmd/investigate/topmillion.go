package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/folbricht/umbrella"
)

const topMillionUsage = `top-million [options]

Show the most popular domains by rank.`

func topMillion(key string, args []string) error {
	var limit int
	flags := flag.NewFlagSet("top-million", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, topMillionUsage)
		flags.PrintDefaults()
	}
	flags.IntVar(&limit, "limit", 500, "Max number of domains")
	flags.Parse(args)
	if flags.NArg() > 0 {
		return tooManyArgs
	}
	opts := umbrella.QueryOptions{
		"limit": {strconv.Itoa(limit)},
	}
	client := umbrella.NewInvestigate(key)
	results, err := client.TopMillion(opts)
	if err != nil {
		return err
	}
	return printJSON(results)
}
