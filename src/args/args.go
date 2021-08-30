package args

import (
	"flag"
	"fmt"
)

// Our global variables
var (
	wordlist string
	output   string
	threads  int
)

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
