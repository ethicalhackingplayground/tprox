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
	Verbose  bool
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

	rand.Seed(time.Now().UnixNano())
	nCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nCPU)

	flag.StringVar(&Wordlist, "w", "", "The wordlist to use against a valid endpoint to traverse")
	flag.StringVar(&Output, "o", "", "Output the results to a file")
	flag.StringVar(&Regex, "r", "", "Filter crawl with regex pattern")
	Crawl := flag.Bool("c", false, "crawl the resolved domain while testing for proxy misconfigs")
	Verbose := flag.Bool("v", false, "Show verbose output")
	flag.IntVar(&Depth, "d", 5, "The crawl depth")
	flag.IntVar(&Threads, "t", 10, "The number of concurrent requests")
	flag.Parse()
	if Wordlist == "" {
		flag.PrintDefaults()
		return false, *Crawl, *Verbose
	} else {
		return true, *Crawl, *Verbose
	}
}
