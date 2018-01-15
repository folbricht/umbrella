package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
)

const usage = `investigate -key <APITOKEN> <command> [options]
investigate <command> -h

Commands:
domain-categories - List category IDs and Labels
domain-categorization - Categorization for a single domain
domain-categorizations - Categorization of multiple domains
domain-timeline - Show the timeline of a domain
domain-volume - Query volume of a domain
search - Perform a pattern search
co-occurrances - Find domains that were queried around the same time
related - Find domains related to a domain
security - Show available security information for a domain
domain-history - Query the history of a domain+type
ip-history - Query the history of an ip+type
`

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
		flag.PrintDefaults()
	}
	var key string
	flag.StringVar(&key, "key", "", "Umbrella API token")
	flag.Parse()

	if len(os.Args) < 2 {
		die(notEnoughArgs)
	}

	cmd := flag.Arg(0)
	args := flag.Args()[1:]
	if key == "" {
		key = os.Getenv("UMBRELLA_KEY")
	}

	handlers := map[string]func(string, []string) error{
		"-h":                     help,
		"domain-categories":      domainCategories,
		"domain-categorization":  domainCategorization,
		"domain-categorizations": domainCategorizations,
		"domain-timeline":        domainTimeline,
		"domain-volume":          domainVolume,
		"search":                 search,
		"co-occurrences":         coOccurrences,
		"related":                related,
		"security":               security,
		"domain-history":         domainHistory,
		"ip-history":             ipHistory,
	}
	h, ok := handlers[cmd]
	if !ok {
		die(fmt.Errorf("Unknown command %s", cmd))
	}

	if err := h(key, args); err != nil {
		die(err)
	}
}

func help(ke string, args []string) error {
	flag.Usage()
	os.Exit(1)
	return nil
}

func die(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func printJSON(v interface{}) error {
	w := json.NewEncoder(os.Stderr)
	w.SetIndent("", "  ")
	return w.Encode(v)
}

var (
	notEnoughArgs = errors.New("Not enough arguments. See -h for help.")
	tooManyArgs   = errors.New("Too many arguments. See -h for help.")
)
