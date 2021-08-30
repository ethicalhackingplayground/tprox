package args

import (
	"flag"
	"math/rand"
	"runtime"
	"time"

	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
)

// Our global variables
var (
	Wordlist string
	Output   string
	Threads  int
	Crawl    bool
	Depth    int
	Regex    string
	Silent   bool
	Scope    string
)

func printBanner() {

	banner := `

 
 	  __                   
	 / /____  _______ __ __
	/ __/ _ \/ __/ _ \\ \ /
	\__/ .__/_/  \___/_\_\   v0.1-dev
   	   /_/                  

	`

	gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)

	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msg("\t\tgithub.com/ethicalhackingplayground\n\n")

	gologger.Info().Msg("Use with caution. You are responsible for your actions\n")
	gologger.Info().Msg("Developers assume no liability and are not responsible for any misuse or damage.\n\n")

}

// Return a true or false if the args are valid.
func ParseArgs() (bool, bool, bool) {

	// Print the banner
	printBanner()

	flag.StringVar(&Wordlist, "w", "", "The wordlist to use against a valid endpoint to traverse")
	flag.StringVar(&Output, "o", "", "Output the results to a file")
	flag.StringVar(&Regex, "regex", "", "Filter crawl with regex pattern")
	flag.StringVar(&Scope, "scope", "", "Specify a scope to crawl in with a regex")
	Crawl := flag.Bool("crawl", false, "crawl the resolved domain while testing for proxy misconfigs")
	Silent := flag.Bool("s", false, "Show Silent output")
	flag.IntVar(&Depth, "depth", 5, "The crawl depth")
	flag.IntVar(&Threads, "c", 10, "The number of concurrent requests")
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	nCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nCPU)

	if Wordlist == "" {
		flag.PrintDefaults()
		return false, *Crawl, *Silent
	} else {
		return true, *Crawl, *Silent
	}
}
