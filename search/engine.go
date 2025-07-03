package search

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var Engines = []string{
	"https://www.google.com/search?q=%s&num=100&start=%d",
	"https://www.bing.com/search?q=%s&count=50&first=%d",
	"https://duckduckgo.com/html/?q=%s&s=%d",
	"https://search.yahoo.com/search?p=%s&n=100&b=%d",
}

var Agents = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
}

func FilterEngines(exclude string) []string {
	if exclude == "" {
		return Engines
	}
	
	excluded := make(map[string]bool)
	for _, e := range strings.Split(exclude, ",") {
		excluded[strings.TrimSpace(e)] = true
	}
	
	var filtered []string
	for _, e := range Engines {
		if !excluded[EngineName(e)] {
			filtered = append(filtered, e)
		}
	}
	return filtered
}

func EngineName(url string) string {
	switch {
	case strings.Contains(url, "google.com"):
		return "google"
	case strings.Contains(url, "bing.com"):
		return "bing"
	case strings.Contains(url, "duckduckgo.com"):
		return "duckduckgo"
	case strings.Contains(url, "yahoo.com"):
		return "yahoo"
	default:
		return "unknown"
	}
}

func Colorize(name string) string {
	switch name {
	case "google":
		return "[\033[34mG\033[31mo\033[33mo\033[34mg\033[32ml\033[31me\033[0m]"
	case "bing":
		return "[\033[38;5;33mbing\033[0m]"
	default:
		return "[" + name + "]"
	}
}

func RandomDelay(min, max int) time.Duration {
	rand.Seed(time.Now().UnixNano())
	return time.Duration(rand.Intn(max-min)+min) * time.Millisecond
}
