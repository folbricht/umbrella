package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/folbricht/umbrella"
)

const domainCategorizationUsage = `domain-categorization [options] <domain>

Categorize a single domain. If 'showlabels' is given, it'll return the string
labels of the categories rather than their number.`

func domainCategorization(key string, args []string) error {
	var (
		showlabels bool
	)
	flags := flag.NewFlagSet("domain-categorization", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, domainCategorizationUsage)
		flags.PrintDefaults()
	}
	flags.BoolVar(&showlabels, "showlabels", false, "show category labels not numbers")
	flags.Parse(args)

	if flags.NArg() < 1 {
		return notEnoughArgs
	}
	if flags.NArg() > 1 {
		return tooManyArgs
	}
	domain := flags.Arg(0)
	client := umbrella.NewInvestigate(key)
	categorization, err := client.DomainCategorization(domain, showlabels)
	if err != nil {
		return err
	}
	return printJSON(categorization)
}

const domainCategorizationsUsage = `domain-categorizations <domain> [<domain>...]

Categorize a multiple domains.`

func domainCategorizations(key string, args []string) error {
	flags := flag.NewFlagSet("domain-categorizations", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, domainCategorizationsUsage)
		flags.PrintDefaults()
	}
	flags.Parse(args)

	if flags.NArg() < 1 {
		return notEnoughArgs
	}
	client := umbrella.NewInvestigate(key)
	domains := flags.Args()
	categorization, err := client.DomainCategorizations(domains...)
	if err != nil {
		return err
	}
	return printJSON(categorization)
}

const domainCategoriesUsage = `domain-categories

List all category numbers and their equivalent labels.`

func domainCategories(key string, args []string) error {
	flags := flag.NewFlagSet("domain-categories", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, domainCategoriesUsage)
		flags.PrintDefaults()
	}
	flags.Parse(args)
	if flags.NArg() > 0 {
		return tooManyArgs
	}
	client := umbrella.NewInvestigate(key)
	categories, err := client.DomainCategories()
	if err != nil {
		return err
	}
	return printJSON(categories)
}

const domainTimelineUsage = `domain-timeline <domain>

Show the timeline for a domain.`

func domainTimeline(key string, args []string) error {
	flags := flag.NewFlagSet("domain-timeline", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, domainTimelineUsage)
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
	timeline, err := client.DomainTimeline(domain)
	if err != nil {
		return err
	}
	return printJSON(timeline)
}

const domainVolumeUsage = `domain-volume [options] <domain>

Shows the query volume for a domain`

func domainVolume(key string, args []string) error {
	var (
		start, stop string
		match       string
	)
	flags := flag.NewFlagSet("domain-volume", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, domainVolumeUsage)
		flags.PrintDefaults()
	}
	flags.StringVar(&start, "start", "-1days", "Start time")
	flags.StringVar(&stop, "stop", "now", "End time")
	flags.StringVar(&match, "match", "all", "Match 'component', 'exact', or 'all'")
	flags.Parse(args)
	if flags.NArg() < 1 {
		return notEnoughArgs
	}
	if flags.NArg() > 1 {
		return tooManyArgs
	}
	domain := flags.Arg(0)
	client := umbrella.NewInvestigate(key)
	opts := umbrella.QueryOptions{
		"start": {start},
		"stop":  {stop},
		"match": {match},
	}
	volume, err := client.DomainVolume(domain, opts)
	if err != nil {
		return err
	}
	return printJSON(volume)
}
