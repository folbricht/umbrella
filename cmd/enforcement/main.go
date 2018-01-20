package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
)

const usage = `enforcement -key <CUSTOMERKEY> <command> [options]
enforcement <command> -h

Commands:
list-domains - List domains currently on the blocklist (includes pagination)
list-all-domains - List all domains currently on the blocklist
delete-domain - Remove a domain from the blocklist
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
		die(errNotEnoughArgs)
	}

	cmd := flag.Arg(0)
	args := flag.Args()[1:]
	if key == "" {
		key = os.Getenv("CUSTOMER_KEY")
	}

	handlers := map[string]func(string, []string) error{
		"-h":               help,
		"list-domains":     getDomains,
		"list-all-domains": getAllDomains,
		"delete-domain":    deleteDomain,
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
	errNotEnoughArgs = errors.New("Not enough arguments. See -h for help.")
	errTooManyArgs   = errors.New("Too many arguments. See -h for help.")
)
