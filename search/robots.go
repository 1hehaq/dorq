package search

import (
	"bufio"
	"net/http"
	"net/url"
	"strings"
)

var robotsCache = make(map[string]bool)
var robotsMu sync.Mutex

func checkRobots(engine string, client *http.Client) bool {
	u, err := url.Parse(engine)
	if err != nil {
		return true
	}
	
	robotsURL := u.Scheme + "://" + u.Host + "/robots.txt"
	robotsMu.Lock()
	defer robotsMu.Unlock()
	
	if allowed, exists := robotsCache[robotsURL]; exists {
		return allowed
	}
	
	resp, err := client.Get(robotsURL)
	if err != nil || resp.StatusCode != 200 {
		robotsCache[robotsURL] = true
		return true
	}
	defer resp.Body.Close()
	
	scanner := bufio.NewScanner(resp.Body)
	disallowed := false
	path := u.Path
	
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "User-agent: *") {
			disallowed = false
		} else if strings.HasPrefix(line, "Disallow:") && !disallowed {
			disallowedPath := strings.TrimSpace(strings.TrimPrefix(line, "Disallow:"))
			if strings.HasPrefix(path, disallowedPath) {
				robotsCache[robotsURL] = false
				return false
			}
		}
	}
	
	robotsCache[robotsURL] = true
	return true
}
