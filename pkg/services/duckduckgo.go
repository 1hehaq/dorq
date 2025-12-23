package services

import (
	"net/url"
	"strings"
	"time"

	"github.com/1hehaq/dorq/pkg/config"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
)

func SearchDuckDuckGo(browser *rod.Browser, query string, config config.Config) ([]string, error) {
	page := browser.MustPage()
	defer page.MustClose()

	serpURL := "https://duckduckgo.com/?ia=web&q=" + url.QueryEscape(query)

	page.Timeout(time.Duration(config.Timeout) * time.Second).MustNavigate(serpURL)
	page.WaitLoad()

	html := page.MustHTML()
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		return nil, err
	}

	var urls []string
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, ok := s.Attr("href")
		if !ok || href == "" {
			return
		}
		if strings.HasPrefix(href, "http") && !strings.Contains(href, "duckduckgo.com") {
			urls = append(urls, href)
		}
	})
	return urls, nil
}
