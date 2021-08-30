package traversal

import (
	"bufio"
	"fmt"
	"github.com/ethicalhackingplayground/tprox/tprox/args"
	"github.com/ethicalhackingplayground/tprox/tprox/discover"
	"github.com/projectdiscovery/gologger"
	"github.com/schollz/progressbar/v3"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
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

			gologger.Debug().Msg("Discovery Interesting Files/Dirs on " + testUrl)
			fmt.Println("")
			// Start bruteforcing for files and directories
			words := make(chan string)

			// Read in the wordlist list
			wordFile, err := os.Open(args.Wordlist)
			count := int64(len(LinesInFile(args.Wordlist)))
			if err != nil {
				return
			}
			wordBytes := bufio.NewScanner(wordFile)
			if err != nil {
				return
			}
			bar := progressbar.Default(count)
			for i := 0; i < args.Threads; i++ {
				wg.Add(1)
				go func() {

					// wordlist brute channel loop
					for word := range words {
						discover.BruteForDirAndFile(client, wg, url, testUrl, word, silent)
						bar.Add(1)
					}
					time.Sleep(40 * time.Millisecond)

				}()
				wg.Done()

			}

			for wordBytes.Scan() {
				words <- wordBytes.Text()
			}
			close(words)

		}
	}
}

func LinesInFile(fileName string) []string {
	f, _ := os.Open(fileName)
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		// Append line to result.
		result = append(result, line)
	}
	return result
}
