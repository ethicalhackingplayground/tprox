package traversal

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/ethicalhackingplayground/tprox/tprox/args"
	"github.com/ethicalhackingplayground/tprox/tprox/discover"
	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
	"io"
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

			white := color.New(color.FgWhite, color.Bold).SprintFunc()
			blue := color.New(color.FgBlue, color.Bold).SprintFunc()

			if !silent {
				fmt.Printf("%s%s%s %s\n", white("["), blue("Proxy"), white("]"), white(testUrl))
			}

			// Start bruteforcing for files and directories
			words := make(chan string)

			// Read in the wordlist list
			wordFile, err := os.Open(args.Wordlist)
			if err != nil {
				return
			}
			wordBytes := bufio.NewScanner(wordFile)
			count, err := lineCounter(wordFile)
			if err != nil {
				gologger.Error().Msg(err.Error())
				return
			}
			bar := progressbar.Default(int64(count))
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

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
