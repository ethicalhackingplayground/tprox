package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/projectdiscovery/gologger"
	"github.com/ethicalhackingplayground/tprox/src/args"
)

// Parse the arguments and run the test function.
func main() {

	if args.parseArgs() {
		gologger.Debug().Msg("[>] Finding misconfigured proxies ")
		fmt.Println("")
		run()
	}
}

// This is where all the path traversal functions begin.
func run() {

	urls := make(chan string)

	var wg sync.WaitGroup
	for i := 0; i < args.threads; i++ {
		wg.Add(1)
		go func() {
			// Url channel loop
			for url := range urls {
				for _, traversal := range args.payloads {
					traversal.testTraversal(&wg, url, traversal)
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
