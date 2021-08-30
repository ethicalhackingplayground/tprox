package traversal

import (
	"bufio"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/projectdiscovery/gologger"
)

// The payloads to test
var payloads = [3]string{"..%2f", "..;/", "../"}

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
			gologger.Info().Msg("[-] Brutforcing proxy " + testUrl)

			// Start bruteforcing for files and directories
			words := make(chan string)

			for i := 0; i < args.threads; i++ {
				wg.Add(1)
				go func() {

					// wordlist brute channel loop
					for word := range words {
						discover.bruteForDirAndFile(client, wg, url, testUrl, word)
					}
				}()
				wg.Done()
			}

			// Read in the wordlist list
			wordFile, err := os.Open(args.wordlist)
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
