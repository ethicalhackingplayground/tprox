package traversal

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/ethicalhackingplayground/tprox/src/args"
	"github.com/ethicalhackingplayground/tprox/src/discover"
	"github.com/projectdiscovery/gologger"
)

// Craft the new url to test, this works as the following.
// 1.) Input the Url to test
// 2.) Calculate the number of paths
// 3.) Append the payload the same number of times as the paths
func CraftTestUrl(count int, url string, payload string) string {
	var traversal = ""
	for i := 0; i < count; i++ {
		traversal = traversal + payload
	}

	return url + "/" + traversal + "/"

}

// Test for proxy traversal attacks
func TestTraversal(wg *sync.WaitGroup, url string, payload string, silent bool) {

	client := http.Client{}

	// Get the test url
	paths := strings.Split(url, "/")
	pathCount := len(paths) - 2
	testUrl := CraftTestUrl(pathCount, url, payload)

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
		testUrl := CraftTestUrl(pathCount, url, payload)

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
			if silent == false {
				gologger.Debug().Msg("Found Proxy, Bruteforcing " + testUrl)
			}

			// Start bruteforcing for files and directories
			words := make(chan string)

			for i := 0; i < args.Threads; i++ {
				wg.Add(1)
				go func() {

					// wordlist brute channel loop
					for word := range words {
						discover.BruteForDirAndFile(client, wg, url, testUrl, word)
					}
				}()
				wg.Done()
			}

			// Read in the wordlist list
			wordFile, err := os.Open(args.Wordlist)
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
