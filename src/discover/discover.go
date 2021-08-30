package discover

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/fatih/color"
)

// Start the content discovery for files and directories
func BruteForDirAndFile(client http.Client, wg *sync.WaitGroup, url string, testUrl string, word string) {

	info := color.New(color.FgWhite, color.Italic).SprintFunc()
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
			fmt.Println("")
			fmt.Printf("%s%s%s %s\n", white("["), green("FOUND"), white("]"), info(contentFound))
		}

	}

}
