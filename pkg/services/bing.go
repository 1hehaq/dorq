package services

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/1hehaq/dorq/pkg/config"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
)

func SearchBing(browser *rod.Browser, query string, config config.Config) ([]string, error) {
	page := browser.MustPage()
	defer page.MustClose()

	serpURL := "https://www.bing.com/search?q=" + url.QueryEscape(query) + "&num=" + strconv.Itoa(config.Limit)

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
		if strings.HasPrefix(href, "http") && !strings.Contains(href, "bing.com") {
			urls = append(urls, href)
		}
	})
	return urls, nil
}
