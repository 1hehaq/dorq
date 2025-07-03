package search

import (
	"crypto/tls"
	"net/http"
	"net/url"
)

func CreateClient(proxy string) *http.Client {
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	
	if proxy != "" {
		if proxyURL, err := url.Parse(proxy); err == nil {
			transport.Proxy = http.ProxyURL(proxyURL)
		}
	}
	
	return &http.Client{
		Transport: transport,
		Timeout:   15 * time.Second,
	}
}

func BuildURL(engine, query string, page int) string {
	escaped := url.QueryEscape(query)
	switch {
	case strings.Contains(engine, "google.com"):
		return fmt.Sprintf(engine, escaped, page*100)
	case strings.Contains(engine, "bing.com"):
		return fmt.Sprintf(engine, escaped, page*50+1)
	default:
		return fmt.Sprintf(engine, escaped, page*30)
	}
}
