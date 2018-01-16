package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/folbricht/umbrella"
)

const latestMaliciousUsage = `latest-malicious <IP>

Find malicious domains associated with an IP.`

func latestMalicious(key string, args []string) error {
	flags := flag.NewFlagSet("latest-malicious", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, latestMaliciousUsage)
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
	results, err := client.LatestMalicious(ip)
	if err != nil {
		return err
	}
	return printJSON(results)
}
