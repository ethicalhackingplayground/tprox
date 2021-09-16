package traversal

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"crypto/tls"

	"github.com/ethicalhackingplayground/tprox/tprox/args"
	"github.com/ethicalhackingplayground/tprox/tprox/discover"
	"github.com/fatih/color"
	"github.com/schollz/progressbar/v3"
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
func TestTraversal(wg *sync.WaitGroup, url string, payload string, silent bool, traverse bool, progress bool, test bool, discoverContent bool, check bool) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	var testUrl = ""

	if test {
		// Get the test url
		paths := strings.Split(url, "/")
		pathCount := len(paths) - 2
		testUrl = CraftTestUrl(pathCount, url, payload)

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
			testUrl = CraftTestUrl(pathCount, url, payload) + "doesnotexist"

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
				color.NoColor = false
				white := color.New(color.FgWhite, color.Bold).SprintFunc()
				blue := color.New(color.FgBlue, color.Bold).SprintFunc()

				fmt.Printf("\n%s%s%s %s\n\n", white("["), blue("PROXY FOUND"), white("]"), white(testUrl))
			}
		}
	} else {
		if check {
			discover.BruteForDirAndFile(*client, wg, url, url, "", silent, traverse, discoverContent, check)
		} else {
			if discoverContent {
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
				if progress {
					bar := progressbar.Default(count)
					for i := 0; i < args.Threads; i++ {
						wg.Add(1)
						go func() {

							// wordlist brute channel loop
							for word := range words {
								discover.BruteForDirAndFile(*client, wg, url+"/", url+"/", word, silent, traverse, discoverContent, check)
								bar.Add(1)
							}
							time.Sleep(40 * time.Millisecond)

						}()
						wg.Done()
					}
				} else {
					for i := 0; i < args.Threads; i++ {
						wg.Add(1)
						go func() {

							// wordlist brute channel loop
							for word := range words {
								discover.BruteForDirAndFile(*client, wg, url, url, word, silent, traverse, discoverContent, check)
							}
							time.Sleep(40 * time.Millisecond)

						}()
						wg.Done()
					}
				}

				for wordBytes.Scan() {
					words <- wordBytes.Text()
				}
				close(words)

			} else {
				if traverse {
					// Get the test url
					paths := strings.Split(url, "/")
					pathCount := len(paths) - 2
					testUrl = CraftTestUrl(pathCount, url, payload)

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
						testUrl = CraftTestUrl(pathCount, url, payload) + "doesnotexist"

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
							color.NoColor = false
							white := color.New(color.FgWhite, color.Bold).SprintFunc()
							blue := color.New(color.FgBlue, color.Bold).SprintFunc()

							fmt.Printf("\n%s%s%s %s\n\n", white("["), blue("PROXY FOUND"), white("]"), white(testUrl))
							if args.Wordlist == "" {
								return
							}
							if !test {
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
								if progress {
									bar := progressbar.Default(count)
									for i := 0; i < args.Threads; i++ {
										wg.Add(1)
										go func() {

											// wordlist brute channel loop
											for word := range words {
												discover.BruteForDirAndFile(*client, wg, url, CraftTestUrl(pathCount, url, payload), word, silent, traverse, discoverContent, check)
												bar.Add(1)
											}
											time.Sleep(40 * time.Millisecond)

										}()
										wg.Done()
									}

								} else {

									for i := 0; i < args.Threads; i++ {
										wg.Add(1)
										go func() {

											// wordlist brute channel loop
											for word := range words {
												discover.BruteForDirAndFile(*client, wg, url, CraftTestUrl(pathCount, url, payload), word, silent, traverse, discoverContent, check)
											}

										}()
										wg.Done()
									}

								}

								for wordBytes.Scan() {
									words <- wordBytes.Text()
								}
								close(words)
							}

						}
					}
				} else {

					testUrl = url + "/"
					color.NoColor = false
					white := color.New(color.FgWhite, color.Bold).SprintFunc()
					blue := color.New(color.FgBlue, color.Bold).SprintFunc()

					fmt.Printf("\n%s%s%s %s\n\n", white("["), blue("DISCOVERY"), white("]"), white(testUrl))
					if args.Wordlist == "" {
						return
					}
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
					if progress {
						bar := progressbar.Default(count)
						for i := 0; i < args.Threads; i++ {
							wg.Add(1)
							go func() {

								// wordlist brute channel loop
								for word := range words {
									discover.BruteForDirAndFile(*client, wg, url, testUrl, word, silent, traverse, discoverContent, check)
									bar.Add(1)
								}
								time.Sleep(40 * time.Millisecond)

							}()
							wg.Done()
						}

					} else {

						for i := 0; i < args.Threads; i++ {
							wg.Add(1)
							go func() {

								// wordlist brute channel loop
								for word := range words {
									discover.BruteForDirAndFile(*client, wg, url, testUrl, word, silent, traverse, discoverContent, check)
								}
							}()

						}
						wg.Done()
					}
					for wordBytes.Scan() {
						words <- wordBytes.Text()
					}
					close(words)
				}
			}
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
