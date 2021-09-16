package discover

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/ethicalhackingplayground/tprox/tprox/args"
	"github.com/fatih/color"
)

// Start the content discovery for files and directories
func BruteForDirAndFile(client http.Client, wg *sync.WaitGroup, url string, testUrl string, word string, silent bool, traverse bool, discover bool, check bool) {
	color.NoColor = false
	if discover {
		info := color.New(color.FgWhite, color.Bold).SprintFunc()
		white := color.New(color.FgWhite, color.Bold).SprintFunc()
		green := color.New(color.FgGreen, color.Bold).SprintFunc()

		contentNotFound := url + "/" + word
		rootDomain := strings.Split(contentNotFound, "/")[0] + "//" + strings.Split(contentNotFound, "/")[2] + "/" + word
		contentFound := testUrl + word
		resp1, err := http.Get(contentFound)
		if err != nil {
			return
		}
		resp2, err := http.Get(rootDomain)
		if err != nil {
			return
		}
		if resp1.StatusCode == 200 && resp2.StatusCode != 200 {
			if args.Output != "" {

				f, err := os.Create(args.Output)
				if err != nil {
					return
				}
				defer f.Close()

				_, err2 := f.WriteString(contentFound + "\n")

				if err2 != nil {
					return
				}
			}
			fmt.Printf("\n\n%s%s%s %s\n\n", white("["), green("FOUND"), white("]"), info(contentFound))
			defer color.Unset() // Use it in your function

		}
	} else {
		if word == "" {
			info := color.New(color.FgWhite, color.Bold).SprintFunc()
			white := color.New(color.FgWhite, color.Bold).SprintFunc()
			green := color.New(color.FgGreen, color.Bold).SprintFunc()

			fullUrl := url
			endPathLength := len(strings.Split(fullUrl, "/"))
			endPath := strings.Split(fullUrl, "/")[endPathLength-1]
			rootDomain := strings.Split(fullUrl, "/")[0] + "//" + strings.Split(fullUrl, "/")[2] + "/" + endPath

			contentFound := testUrl + word
			resp1, err := http.Get(contentFound)
			if err != nil {
				return
			}

			resp2, err := http.Get(rootDomain)
			if err != nil {
				return
			}

			if traverse {
				if resp1.StatusCode == 200 && resp2.StatusCode != 200 {
					if args.Output != "" {

						f, err := os.Create(args.Output)
						if err != nil {
							return
						}
						defer f.Close()

						_, err2 := f.WriteString(contentFound + "\n")

						if err2 != nil {
							return
						}
					}
					fmt.Printf("\n\n%s%s%s %s\n\n", white("["), green("FOUND"), white("]"), info(contentFound))
					defer color.Unset() // Use it in your function

				}
			} else {

				if resp1.StatusCode == 200 {
					if args.Output != "" {

						f, err := os.Create(args.Output)
						if err != nil {
							return
						}
						defer f.Close()

						_, err2 := f.WriteString(contentFound + "\n")

						if err2 != nil {
							return
						}
					}

					fmt.Printf("\n\n%s%s%s %s\n\n", white("["), green("FOUND"), white("]"), info(contentFound))
					defer color.Unset() // Use it in your function
				}
			}
		} else {
			info := color.New(color.FgWhite, color.Bold).SprintFunc()
			white := color.New(color.FgWhite, color.Bold).SprintFunc()
			green := color.New(color.FgGreen, color.Bold).SprintFunc()

			contentNotFound := url + "/" + word
			rootDomain := strings.Split(contentNotFound, "/")[0] + "//" + strings.Split(contentNotFound, "/")[2] + "/" + word
			contentFound := testUrl + word
			resp1, err := http.Get(contentFound)
			if err != nil {
				return
			}
			resp2, err := http.Get(rootDomain)
			if err != nil {
				return
			}

			if traverse {
				if resp1.StatusCode == 200 && resp2.StatusCode != 200 {
					if args.Output != "" {

						f, err := os.Create(args.Output)
						if err != nil {
							return
						}
						defer f.Close()

						_, err2 := f.WriteString(contentFound + "\n")

						if err2 != nil {
							return
						}
					}
					fmt.Printf("\n\n%s%s%s %s\n\n", white("["), green("FOUND"), white("]"), info(contentFound))
					defer color.Unset() // Use it in your function

				}
			} else {

				if resp1.StatusCode == 200 {
					if args.Output != "" {

						f, err := os.Create(args.Output)
						if err != nil {
							return
						}
						defer f.Close()

						_, err2 := f.WriteString(contentFound + "\n")

						if err2 != nil {
							return
						}
					}

					fmt.Printf("\n\n%s%s%s %s\n\n", white("["), green("FOUND"), white("]"), info(contentFound))
					defer color.Unset() // Use it in your function
				}
			}
		}

	}
}
