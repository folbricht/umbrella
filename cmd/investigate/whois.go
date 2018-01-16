package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/folbricht/umbrella"
)

const whoisEmailUsage = `whois-email [options] <email>

Find domains registered with an email.`

func whoisEmail(key string, args []string) error {
	var limit int
	flags := flag.NewFlagSet("whois-email", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, whoisEmailUsage)
		flags.PrintDefaults()
	}
	flags.IntVar(&limit, "limit", 500, "Max number of domains")
	flags.Parse(args)
	if flags.NArg() < 1 {
		return notEnoughArgs
	}
	if flags.NArg() > 1 {
		return tooManyArgs
	}
	email := flags.Arg(0)
	client := umbrella.NewInvestigate(key)
	results, err := client.WHOISEmail(limit, email)
	if err != nil {
		return err
	}
	return printJSON(results)
}

const whoisEmailsUsage = `whois-emails [options] <email> [<email>...]

Find domains registered for multiple emails.`

func whoisEmails(key string, args []string) error {
	var limit int
	flags := flag.NewFlagSet("whois-email", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, whoisEmailsUsage)
		flags.PrintDefaults()
	}
	flags.IntVar(&limit, "limit", 500, "Max number of domains per email")
	flags.Parse(args)
	if flags.NArg() < 1 {
		return notEnoughArgs
	}
	if flags.NArg() > 1 {
		return tooManyArgs
	}
	emails := flags.Args()
	client := umbrella.NewInvestigate(key)
	results, err := client.WHOISEmails(limit, emails...)
	if err != nil {
		return err
	}
	return printJSON(results)
}
