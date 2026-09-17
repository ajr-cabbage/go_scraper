package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	_, ok := cfg.pages[normalizedURL]
	if ok {
		return false
	}
	return true
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()
	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()
	parsedCurrentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Print(err)
		return
	}
	// check if url is in the base domain
	if cfg.baseURL.Host != parsedCurrentURL.Host {
		return
	}
	// normalize rawCurrentURL and check if we visited
	normUrl, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Print(err)
	}
	// stop if this isn't the first visit
	if !cfg.addPageVisit(normUrl) {
		fmt.Println(" - Been here already")
		return
	}
	// otherwise proceed with crawl
	fmt.Printf(" - Crawling: %s\n", normUrl)
	html, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Print(err)
		return
	}
	// add page data to pages map
	pageDat := extractPageData(html, rawCurrentURL)
	cfg.mu.Lock()
	cfg.pages[normUrl] = pageDat
	cfg.mu.Unlock()
	for _, link := range pageDat.OutgoingLinks {
		//fmt.Printf(" - Crawling Page: %s\n", link)
		cfg.wg.Add(1)
		go cfg.crawlPage(link)
	}
}
