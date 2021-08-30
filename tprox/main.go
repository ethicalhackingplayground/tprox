package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sync"

	"github.com/ethicalhackingplayground/tprox/src/args"
	"github.com/ethicalhackingplayground/tprox/src/traversal"
	"github.com/gocolly/colly"
	"github.com/projectdiscovery/gologger"
)

// The payloads to test
var Payloads = [3]string{"..%2f", "..;/", "../"}

// Parse the arguments and run the test function.
func main() {
	parsed, crawl, silent := args.ParseArgs()
	if parsed {
		if silent == false {
			gologger.Debug().Msg("Finding misconfigured proxies ")
			fmt.Println("")
		}

		run(crawl, silent)
	}
}

// This is where all the path traversal functions begin.
func run(crawl bool, silent bool) {

	urls := make(chan string)
	// Crawling is enabled
	c := colly.NewCollector(
		colly.MaxDepth(args.Depth),
	)
	var wg sync.WaitGroup
	for i := 0; i < args.Threads; i++ {
		wg.Add(1)
		go func() {
			// Url channel loop
			for url := range urls {
				for _, p := range Payloads {
					if crawl {
						Crawl(c, &wg, url, p, silent)

					} else {
						traversal.TestTraversal(&wg, url, p, silent)
					}
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
		if silent == false {
			gologger.Debug().Msg("Crawled " + r.URL.String())
		}
		matched, _ := regexp.MatchString(args.Regex, r.URL.String())
		if args.Regex != "" {
			if matched {
				traversal.TestTraversal(wg, r.URL.String(), payload, silent)
			}
		} else {
			traversal.TestTraversal(wg, r.URL.String(), payload, silent)
		}

	})
	c.Visit(url)
}