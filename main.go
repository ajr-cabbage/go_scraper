package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	args := os.Args[1:]

	switch len(args) {
	case 0:
		fmt.Println("usage: scraper <url> <max-concurrency> <max pages>")
		os.Exit(1)
	case 1:
		fmt.Println("usage: scraper <url> <max-concurrency> <max pages>")
		os.Exit(1)
	case 2:
		fmt.Println("usage: scraper <url> <max-concurrency> <max pages>")
		os.Exit(1)
	case 3:
		startingURL := args[0]
		maxConcurrency, err := strconv.Atoi(args[1])
		if err != nil {
			log.Fatal(err)
		}
		maxPages, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("starting crawl of: %s\n", args[0])
		startURL, err := url.Parse(args[0])
		if err != nil {
			log.Fatalf("Bad Starting URL: %s\n", args[0])
		}
		cfg := &config{
			pages:              make(map[string]PageData),
			baseURL:            startURL,
			mu:                 &sync.Mutex{},
			concurrencyControl: make(chan struct{}, maxConcurrency),
			wg:                 &sync.WaitGroup{},
			maxPages:           maxPages,
		}
		cfg.wg.Add(1)
		go cfg.crawlPage(startingURL)
		cfg.wg.Wait()
		for pgURL, pgDat := range cfg.pages {
			fmt.Printf("Page: %s    Heading: %s\n", pgURL, pgDat.Heading)
		}
		err = writeJSONReport(cfg.pages, "report.json")
		if err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	fmt.Printf("Completed in %s\n", time.Since(start))
}
