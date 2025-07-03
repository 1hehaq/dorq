package search

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type ClientPool struct {
	clients map[string]*http.Client
	limiters map[string]*rate.Limiter
	mu sync.Mutex
}

func NewClientPool(proxy string) *ClientPool {
	return &ClientPool{
		clients: make(map[string]*http.Client),
		limiters: make(map[string]*rate.Limiter),
	}
}

func (p *ClientPool) Get(domain string, proxy string) (*http.Client, *rate.Limiter) {
	p.mu.Lock()
	defer p.mu.Unlock()
	
	if client, exists := p.clients[domain]; exists {
		return client, p.limiters[domain]
	}
	
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	
	if proxy != "" {
		if proxyURL, err := url.Parse(proxy); err == nil {
			transport.Proxy = http.ProxyURL(proxyURL)
		}
	}
	
	client := &http.Client{
		Transport: transport,
		Timeout:   15 * time.Second,
	}
	
	limiter := rate.NewLimiter(rate.Every(time.Second), 1)
	p.clients[domain] = client
	p.limiters[domain] = limiter
	
	return client, limiter
}
