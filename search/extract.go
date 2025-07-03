package dorq

import (
	"io"
	"net/http"
	"regexp"
	"strings"
)

func Run(query string, engines []string, proxy string, verbose bool) {
	client := CreateClient(proxy)
	seen := make(map[string]bool)
	
	for _, engine := range engines {
		for page := 0; page < 10; page++ {
			url := BuildURL(engine, query, page)
			req, _ := http.NewRequest("GET", url, nil)
			req.Header.Set("User-Agent", Agents[rand.Intn(len(Agents))])
			
			resp, err := client.Do(req)
			if err != nil || resp.StatusCode != 200 {
				continue
			}
			
			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			
			for _, link := range ExtractLinks(string(body)) {
				if !seen[link] {
					if verbose {
						println(Colorize(EngineName(engine)), link)
					} else {
						println(link)
					}
					seen[link] = true
				}
			}
			
			time.Sleep(RandomDelay(500, 1500))
		}
	}
}

func ExtractLinks(html string) []string {
	re := regexp.MustCompile(`href="(https?://[^"]+)"`)
	matches := re.FindAllStringSubmatch(html, -1)
	var links []string
	
	for _, m := range matches {
		link := m[1]
		if !strings.Contains(link, "google.com") &&
		   !strings.Contains(link, "bing.com") &&
		   !strings.Contains(link, "duckduckgo.com") {
			links = append(links, link)
		}
	}
	return links
}
