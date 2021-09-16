package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"runtime"
	"sync"
	"time"

	"github.com/ethicalhackingplayground/tprox/tprox/args"
	"github.com/ethicalhackingplayground/tprox/tprox/traversal"
	"github.com/gocolly/colly"
	"github.com/projectdiscovery/gologger"
)

// The payloads to test
var Payloads = [15]string{
	"..%2f",
	"..;",
	"%2e%2e%2f",
	"..",
	"..;a=a",
	"%2e%2e%3b",
	"..%00",
	"..%0d",
	"..%5c",
	"..%ff",
	".%2e",
	"%3f",
	"%26",
	"%23",
}
var link = ""
var links = []string{}

// Parse the arguments and run the test function.
func main() {
	parsed, crawl, silent, traverse, progress, test := args.ParseArgs()
	if parsed {
		gologger.Debug().Msg("Finding misconfigured proxies")
		fmt.Println("")
		rand.Seed(time.Now().UnixNano())
		nCPU := runtime.NumCPU()
		runtime.GOMAXPROCS(nCPU)
		run(crawl, silent, traverse, progress, test)
	}
}

// This is where all the path traversal functions begin.
func run(crawl bool, silent bool, traverse bool, progress bool, test bool) {

	urls := make(chan string)

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
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 10})

	wg := sync.WaitGroup{}
	if crawl {
		uscanner := bufio.NewScanner(os.Stdin)
		for uscanner.Scan() {
			link = uscanner.Text()

		}
		for i := 0; i <= args.Threads; i++ {
			wg.Add(1)
			go Crawl(urls, &wg, c, link, silent)

		}
		if !traverse {
			for u := range urls {
				// Url channel loop
				if !silent {
					gologger.Debug().Msg("Crawled " + u)
				}
				traversal.TestTraversal(&wg, u, "", silent, traverse, progress, test)
			}
		} else {
			for _, p := range Payloads {

				for u := range urls {
					// Url channel loop
					if !silent {
						gologger.Debug().Msg("Crawled " + u)
					}
					traversal.TestTraversal(&wg, u, p, silent, traverse, progress, test)
				}

			}
		}

		wg.Wait()

	} else {

		for i := 0; i <= args.Threads; i++ {
			wg.Add(1)
			go func() {

				// Url channel loop
				for url := range urls {

					if traverse == false {
						traversal.TestTraversal(&wg, url, "", silent, traverse, progress, test)
					} else {
						for _, p := range Payloads {
							traversal.TestTraversal(&wg, url, p, silent, traverse, progress, test)
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

}

// Crawl the host
func Crawl(links chan string, wg *sync.WaitGroup, c *colly.Collector, url string, silent bool) {

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		e.Request.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		match, _ := regexp.MatchString(args.Regex, r.URL.String())
		inScope, _ := regexp.MatchString(args.Scope, r.URL.String())

		if args.Regex != "" && args.Scope != "" && match && inScope {

			links <- r.URL.String()

		} else {
			if args.Regex != "" {
				if match {

					links <- r.URL.String()

				}

			} else if args.Scope != "" {
				if inScope {

					links <- r.URL.String()

				}

			} else {

				links <- r.URL.String()
			}

		}
	})

	// Start scraping on https://en.wikipedia.org
	c.Visit(url)
	wg.Done()
	c.Wait()

}
