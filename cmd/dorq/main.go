package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/charmbracelet/log"

	"github.com/1hehaq/dorq/pkg/dorq"
)

const version = "0.1.0"

func init() {
	log.SetTimeFormat("15:04:05")
	log.SetLevel(log.DebugLevel)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			os.Exit(0)
		}
	}()

	query, limit, jsonOutput, showHelp := parseFlags()

	if showHelp {
		displayHelp()
		return
	}

	results, err := searchGoogleDork(query, limit)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	if jsonOutput {
		json.NewEncoder(os.Stdout).Encode(results)
	} else {
		for _, item := range results {
			fmt.Println(item)
		}
	}
}

func parseFlags() (string, int, bool, bool) {
	query := flag.String("q", "", "search query (required)")
	limit := flag.Int("n", 10, "max results")
	jsonOutput := flag.Bool("json", false, "stdout in JSON format")
	showHelp := flag.Bool("h", false, "show help")
	showVersion := flag.Bool("v", false, "show version")
	update := flag.Bool("up", false, "update to latest version")
	flag.Parse()

	if *showVersion {
		displayVersion()
		os.Exit(0)
	}

	if *update {
		performUpdate()
		os.Exit(0)
	}

	if *showHelp {
		return "", 0, false, true
	}

	if *query == "" {
		displayHelp()
		os.Exit(0)
	}

	if *limit <= 0 {
		*limit = 10
	}

	return *query, *limit, *jsonOutput, false
}

func searchGoogleDork(query string, limit int) ([]string, error) {
	results, err := dorq.SearchDork(query, limit)
	if err != nil {
		return nil, err
	}

	return results, nil
}
