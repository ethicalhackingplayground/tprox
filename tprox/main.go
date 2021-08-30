package main

import (
	"bufio"
	"fmt"
	"github.com/ethicalhackingplayground/tprox/tprox/args"
	"github.com/ethicalhackingplayground/tprox/tprox/traversal"
	"github.com/gocolly/colly"
	"github.com/projectdiscovery/gologger"
	"os"
	"regexp"
	"sync"
)

// The payloads to test
var Payloads = [3]string{"..%2f", "..;/", "%2e%2e%2f"}

var crawledUrls = []string{}

// Parse the arguments and run the test function.
func main() {
	parsed, crawl, silent := args.ParseArgs()
	if parsed {
		gologger.Info().Msg("Finding misconfigured proxies")
		fmt.Println("")
		run(crawl, silent)
	}
}

// This is where all the path traversal functions begin.
func run(crawl bool, silent bool) {

	urls := make(chan string)
	crawls := make(chan string)

	// Create a new crolly collector
	c := colly.NewCollector(
		colly.MaxDepth(args.Depth),
		colly.Async(true),
	)

	// Limit the maximum parallelism to 2
	// This is necessary if the goroutines are dynamically
	// created to control the limit of simultaneous requests.
	//
	// Parallelism can be controlled also by spawning fixed
	// number of go routines.
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})

	var wg sync.WaitGroup
	for i := 0; i < args.Threads; i++ {
		wg.Add(1)
		go func() {

			// Url channel loop
			for url := range urls {
				for _, p := range Payloads {
					if crawl {
						Crawl(c, url, silent)
						for urlCrawled := range crawls {
							traversal.TestTraversal(wg, urlCrawled, p, silent)
						}

					} else {
						traversal.TestTraversal(&wg, url, p, silent)
					}
				}

			}
			wg.Done()
		}()

	}

	for _, crawledUrl := range crawledUrls {
		crawls <- crawledUrl
	}

	uscanner := bufio.NewScanner(os.Stdin)
	for uscanner.Scan() {
		urls <- uscanner.Text()
	}

	close(crawls)
	close(urls)
	wg.Wait()
}

// Crawl the host
func Crawl(c *colly.Collector, url string, silent bool) {

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		e.Request.Visit(e.Request.AbsoluteURL(link))
	})

	// The request of each link visisted
	c.OnRequest(func(r *colly.Request) {

		match, _ := regexp.MatchString(args.Regex, r.URL.String())
		inScope, _ := regexp.MatchString(args.Scope, r.URL.String())

		if args.Regex != "" && args.Scope != "" && match && inScope {

			if !silent {
				gologger.Debug().Msg("Crawled " + r.URL.String())
			}
			crawledUrls = append(crawledUrls, r.URL.String())

		} else {
			if args.Regex != "" {
				if match {
					if !silent {
						gologger.Debug().Msg("Crawled " + r.URL.String())
					}
					crawledUrls = append(crawledUrls, r.URL.String())
				}

			} else if args.Scope != "" {
				if inScope {
					if !silent {
						gologger.Debug().Msg("Crawled " + r.URL.String())
					}
					crawledUrls = append(crawledUrls, r.URL.String())
				}

			} else {
				if !silent {
					gologger.Debug().Msg("Crawled " + r.URL.String())
				}
				crawledUrls = append(crawledUrls, r.URL.String())
			}

		}

	})
	c.Visit(url)
	// Wait until threads are finished
	c.Wait()
}
