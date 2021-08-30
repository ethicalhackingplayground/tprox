package discover

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/projectdiscovery/gologger"
)

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
			gologger.Debug().Msg("[*] Found something internal " + contentFound)
		}

	}

}
