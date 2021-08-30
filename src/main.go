package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"

	"github.com/ethicalhackingplayground/tprox/src/args"
	"github.com/ethicalhackingplayground/tprox/src/traversal"
	"github.com/projectdiscovery/gologger"
)

// The payloads to test
var Payloads = [3]string{"..%2f", "..;/", "../"}

// Parse the arguments and run the test function.
func main() {

	if args.ParseArgs() {
		gologger.Debug().Msg("[>] Finding misconfigured proxies ")
		fmt.Println("")
		run()
	}
}

// This is where all the path traversal functions begin.
func run() {

	urls := make(chan string)

	var wg sync.WaitGroup
	for i := 0; i < args.Threads; i++ {
		wg.Add(1)
		go func() {
			// Url channel loop
			for url := range urls {
				for _, p := range Payloads {
					traversal.TestTraversal(&wg, url, p)
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
