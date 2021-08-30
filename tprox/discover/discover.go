package discover

import (
	"fmt"
	"github.com/ethicalhackingplayground/tprox/tprox/args"
	"github.com/fatih/color"
	"net/http"
	"os"
	"sync"
)

// Start the content discovery for files and directories
func BruteForDirAndFile(client http.Client, wg *sync.WaitGroup, url string, testUrl string, word string, silent bool) {
	color.NoColor = false
	info := color.New(color.FgWhite, color.Bold).SprintFunc()
	white := color.New(color.FgWhite, color.Bold).SprintFunc()
	green := color.New(color.FgGreen, color.Bold).SprintFunc()

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
