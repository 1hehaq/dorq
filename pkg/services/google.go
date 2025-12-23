package services

import (
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/1hehaq/dorq/pkg/dorq"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
)

func SearchGoogle(browser *rod.Browser, query string, config dorq.Config) ([]string, error) {
	page := browser.MustPage()
	defer page.MustClose()

	values := url.Values{}
	values.Set("q", query)
	values.Set("num", strconv.Itoa(config.Limit))
	values.Set("ie", "UTF-8")
	values.Set("gbv", "1")
	serpURL := "https://www.google.com/search?" + values.Encode()

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
		if strings.HasPrefix(href, "/url?q=") {
			u, err := url.Parse(href)
			if err == nil {
				q := u.Query().Get("q")
				if q != "" {
					href = q
				}
			}
		}
		if strings.HasPrefix(href, "http") && !strings.Contains(href, "google.com") {
			urls = append(urls, href)
		}
	})
	return urls, nil
}
