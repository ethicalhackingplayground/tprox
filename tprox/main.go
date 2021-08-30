package main

import (
	"bufio"
	"github.com/ethicalhackingplayground/tprox/tprox/args"
	"github.com/ethicalhackingplayground/tprox/tprox/traversal"
	"github.com/gocolly/colly"
	"github.com/projectdiscovery/gologger"
	"os"
	"regexp"
	"sync"
)

// The payloads to test
var Payloads = [3]string{"..%2f", "..;/", "../"}

// Parse the arguments and run the test function.
func main() {
	parsed, crawl, silent := args.ParseArgs()
	if parsed {
		gologger.Debug().Msg("Finding misconfigured proxies ")

		run(crawl, silent)
	}
}

// This is where all the path traversal functions begin.
func run(crawl bool, silent bool) {

	currPayload := ""

	urls := make(chan string)
	// Crawling is enabled
	c := colly.NewCollector(
		// Visit only these root domain
		colly.MaxDepth(args.Depth),
	)

	var wg sync.WaitGroup
	for i := 0; i < args.Threads; i++ {
		wg.Add(1)
		go func() {
			for _, p := range Payloads {
				currPayload = p
			}
			// Url channel loop
			for url := range urls {
				if crawl {
					Crawl(c, &wg, url, currPayload, silent)

				} else {
					traversal.TestTraversal(&wg, url, currPayload, silent)
				}

			}
			wg.Done()
		}()

	}

	uscanner := bufio.NewScanner(os.Stdin)
	for uscanner.Scan() {
		urls <- uscanner.Text()
	}

	close(urls)
	wg.Wait()
}

// Crawl the host
func Crawl(c *colly.Collector, wg *sync.WaitGroup, url string, payload string, silent bool) {

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
			traversal.TestTraversal(wg, r.URL.String(), payload, silent)

		} else {
			if args.Regex != "" {
				if match {
					if !silent {
						gologger.Debug().Msg("Crawled " + r.URL.String())
					}
					traversal.TestTraversal(wg, r.URL.String(), payload, silent)
				}

			} else if args.Scope != "" {
				if inScope {
					if !silent {
						gologger.Debug().Msg("Crawled " + r.URL.String())
					}
					traversal.TestTraversal(wg, r.URL.String(), payload, silent)
				}

			} else {
				if !silent {
					gologger.Debug().Msg("Crawled " + r.URL.String())
				}
				traversal.TestTraversal(wg, r.URL.String(), payload, silent)
			}

		}

	})
	c.Visit(url)
}
