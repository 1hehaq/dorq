package search

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

func Run(query string, engines []string, proxy string, verbose bool) {
	pool := NewClientPool(proxy)
	seen := make(map[string]bool)
	
	for _, engine := range engines {
		domain := getDomain(engine)
		client, limiter := pool.Get(domain, proxy)
		
		for page := 0; page < 10; page++ {
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			if err := limiter.Wait(ctx); err != nil {
				cancel()
				continue
			}
			cancel()
			
			if !checkRobots(engine, client) {
				continue
			}
			
			url := BuildURL(engine, query, page)
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("User-Agent", Agents[rand.Intn(len(Agents))])
			
			resp, err := client.Do(req)
			if err != nil || resp.StatusCode != 200 {
				if resp != nil {
					resp.Body.Close()
				}
				continue
			}
			
			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			
			for _, link := range extractLinks(string(body)) {
				if !seen[link] {
					if verbose {
						println(Colorize(EngineName(engine)), link)
					} else {
						println(link)
					}
					seen[link] = true
				}
			}
			
			time.Sleep(time.Duration(rand.Intn(500)+200) * time.Millisecond)
		}
	}
}

func extractLinks(html string) []string {
	re := regexp.MustCompile(`href="(https?://[^"]+)"`)
	matches := re.FindAllStringSubmatch(html, -1)
	var links []string
	
	for _, m := range matches {
		link := m[1]
		if !strings.Contains(link, "google.com") &&
		   !strings.Contains(link, "bing.com") &&
		   !strings.Contains(link, "duckduckgo.com") &&
		   !strings.Contains(link, "yahoo.com") {
			links = append(links, link)
		}
	}
	return links
}

func getDomain(rawURL string) string {
	u, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	return u.Hostname()
}
