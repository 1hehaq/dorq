package dorq

import (
	"fmt"
	"sync"

	"github.com/1hehaq/dorq/pkg/services"
	"github.com/corpix/uarand"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

var sources = []string{"google", "bing", "duckduckgo"}

func SearchDork(query string, limit int) ([]string, error) {
	ua := uarand.GetRandom()
	browser := rod.New().ControlURL(launcher.New().Set("user-agent", ua).MustLaunch()).MustConnect()
	defer browser.MustClose()

	var wg sync.WaitGroup
	urlChan := make(chan []string, len(sources))

	for _, source := range sources {
		wg.Add(1)
		go func(src string) {
			defer wg.Done()
			urls, err := searchSource(browser, src, query, limit)
			if err == nil {
				urlChan <- urls
			} else {
				urlChan <- []string{}
			}
		}(source)
	}

	go func() {
		wg.Wait()
		close(urlChan)
	}()

	var allUrls []string
	for urls := range urlChan {
		allUrls = append(allUrls, urls...)
	}

	// dedup
	seen := make(map[string]bool)
	var unique []string
	for _, u := range allUrls {
		if !seen[u] {
			seen[u] = true
			unique = append(unique, u)
		}
	}

	if len(unique) > limit {
		unique = unique[:limit]
	}

	if len(unique) == 0 {
		return nil, fmt.Errorf("no_results")
	}

	return unique, nil
}

func searchSource(browser *rod.Browser, source, query string, limit int) ([]string, error) {
	switch source {
	case "google":
		return services.SearchGoogle(browser, query, limit)
	case "bing":
		return services.SearchBing(browser, query, limit)
	case "duckduckgo":
		return services.SearchDuckDuckGo(browser, query, limit)
	default:
		return nil, nil
	}
}
