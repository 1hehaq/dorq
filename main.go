package main

import (
	"bufio"
	"flag"
	"os"
	"strings"

	"github.com/1hehaq/dorq/search"
)

var (
	query   = flag.String("q", "", "dork query or file path")
	verbose = flag.Bool("v", false, "verbose output")
	exclude = flag.String("x", "", "exclude engines (comma-separated)")
	proxy   = flag.String("proxy", "", "proxy URL")
)

func main() {
	flag.Parse()
	engines := search.FilterEngines(*exclude)
	queries := getQueries()
	
	for _, q := range queries {
		search.Run(q, engines, *proxy, *verbose)
	}
}

func getQueries() []string {
	if *query == "" {
		scanner := bufio.NewScanner(os.Stdin)
		var queries []string
		for scanner.Scan() {
			if q := strings.TrimSpace(scanner.Text()); q != "" {
				queries = append(queries, q)
			}
		}
		return queries
	}
	
	if file, err := os.Open(*query); err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		var queries []string
		for scanner.Scan() {
			if q := strings.TrimSpace(scanner.Text()); q != "" {
				queries = append(queries, q)
			}
		}
		return queries
	}
	
	return []string{*query}
}
