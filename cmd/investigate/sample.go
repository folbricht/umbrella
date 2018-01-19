package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/folbricht/umbrella"
)

const samplesUsage = `samples [options] <IP|domain|URL>

List samples associated with an IP, domain, or URL.`

func samples(key string, args []string) error {
	var (
		limit  int
		offset int
		sortby string
	)
	flags := flag.NewFlagSet("samples", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, samplesUsage)
		flags.PrintDefaults()
	}
	flags.IntVar(&limit, "limit", 10, "Max number of samples in the response")
	flags.IntVar(&offset, "offset", 0, "Starting position, used for pagination")
	flags.StringVar(&sortby, "sortby", "score", "Sorting by 'first-seen', 'last-seen', or 'score'")
	flags.Parse(args)
	if flags.NArg() < 1 {
		return notEnoughArgs
	}
	if flags.NArg() > 1 {
		return tooManyArgs
	}
	opts := umbrella.QueryOptions{
		"limit":  {strconv.Itoa(limit)},
		"offset": {strconv.Itoa(offset)},
		"sortby": {sortby},
	}
	term := flags.Arg(0)
	client := umbrella.NewInvestigate(key)
	results, err := client.GetSamples(term, opts)
	if err != nil {
		return err
	}
	return printJSON(results)
}

const sampleUsage = `sample [options] <hash>

Show information about a sample.`

func sample(key string, args []string) error {
	var (
		limit  int
		offset int
	)
	flags := flag.NewFlagSet("sample", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, sampleUsage)
		flags.PrintDefaults()
	}
	flags.IntVar(&limit, "limit", 10, "Max number of samples in the response")
	flags.IntVar(&offset, "offset", 0, "Starting position, used for pagination")
	flags.Parse(args)
	if flags.NArg() < 1 {
		return notEnoughArgs
	}
	if flags.NArg() > 1 {
		return tooManyArgs
	}
	opts := umbrella.QueryOptions{
		"limit":  {strconv.Itoa(limit)},
		"offset": {strconv.Itoa(offset)},
	}
	hash := flags.Arg(0)
	client := umbrella.NewInvestigate(key)
	results, err := client.GetSample(hash, opts)
	if err != nil {
		return err
	}
	return printJSON(results)
}

const sampleArtifactsUsage = `sample-artifacts [options] <hash>

Show information about artifacts associated with a sample.`

func sampleArtifacts(key string, args []string) error {
	var (
		limit  int
		offset int
	)
	flags := flag.NewFlagSet("sample-artifacts", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, sampleArtifactsUsage)
		flags.PrintDefaults()
	}
	flags.IntVar(&limit, "limit", 10, "Max number of samples in the response")
	flags.IntVar(&offset, "offset", 0, "Starting position, used for pagination")
	flags.Parse(args)
	if flags.NArg() < 1 {
		return notEnoughArgs
	}
	if flags.NArg() > 1 {
		return tooManyArgs
	}
	opts := umbrella.QueryOptions{
		"limit":  {strconv.Itoa(limit)},
		"offset": {strconv.Itoa(offset)},
	}
	hash := flags.Arg(0)
	client := umbrella.NewInvestigate(key)
	results, err := client.GetSampleArtifacts(hash, opts)
	if err != nil {
		return err
	}
	return printJSON(results)
}

const sampleConnectionsUsage = `sample-connections [options] <hash>

Show information about connections associated with a sample.`

func sampleConnections(key string, args []string) error {
	var (
		limit  int
		offset int
	)
	flags := flag.NewFlagSet("sample-connections", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, sampleConnectionsUsage)
		flags.PrintDefaults()
	}
	flags.IntVar(&limit, "limit", 10, "Max number of samples in the response")
	flags.IntVar(&offset, "offset", 0, "Starting position, used for pagination")
	flags.Parse(args)
	if flags.NArg() < 1 {
		return notEnoughArgs
	}
	if flags.NArg() > 1 {
		return tooManyArgs
	}
	opts := umbrella.QueryOptions{
		"limit":  {strconv.Itoa(limit)},
		"offset": {strconv.Itoa(offset)},
	}
	hash := flags.Arg(0)
	client := umbrella.NewInvestigate(key)
	results, err := client.GetSampleConnections(hash, opts)
	if err != nil {
		return err
	}
	return printJSON(results)
}

const sampleBehaviorsUsage = `sample-behaviors <hash>

Show information about indicators associated with a sample.`

func sampleBehaviors(key string, args []string) error {
	flags := flag.NewFlagSet("sample-behaviors", flag.ExitOnError)
	flags.Usage = func() {
		fmt.Fprintln(os.Stderr, sampleBehaviorsUsage)
		flags.PrintDefaults()
	}
	flags.Parse(args)
	if flags.NArg() < 1 {
		return notEnoughArgs
	}
	if flags.NArg() > 1 {
		return tooManyArgs
	}
	hash := flags.Arg(0)
	client := umbrella.NewInvestigate(key)
	results, err := client.GetSampleBehaviors(hash)
	if err != nil {
		return err
	}
	return printJSON(results)
}
