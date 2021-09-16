package args

import (
	"flag"

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
	Traverse bool
	Progress bool
	Test     bool
	Check    bool
	Discover bool
)

func printBanner() {

	banner := `

 
 	  __                   
	 / /____  _______ __ __
	/ __/ _ \/ __/ _ \\ \ /
	\__/ .__/_/  \___/_\_\   v0.2-dev
   	   /_/                  

	`

	gologger.DefaultLogger.SetMaxLevel(levels.LevelDebug)

	gologger.Print().Msgf("%s\n", banner)
	gologger.Print().Msg("\t\tgithub.com/ethicalhackingplayground\n\n")

	gologger.Info().Msg("Use with caution. You are responsible for your actions\n")
	gologger.Info().Msg("Developers assume no liability and are not responsible for any misuse or damage.\n\n")

}

// Return a true or false if the args are valid.
func ParseArgs() (bool, bool, bool, bool, bool, bool, bool, bool) {

	// Print the banner
	printBanner()

	flag.StringVar(&Wordlist, "w", "", "The wordlist to use against a valid endpoint to traverse")
	flag.StringVar(&Output, "o", "", "Output the results to a file")
	flag.StringVar(&Regex, "regex", "", "Filter crawl with regex pattern")
	flag.StringVar(&Scope, "scope", "", "Specify a scope to crawl with in using regexs")
	Crawl := flag.Bool("crawl", false, "crawl the resolved domain while testing for proxy misconfigs")
	Silent := flag.Bool("silent", false, "Show Silent output")
	Traverse := flag.Bool("traverse", false, "This flag will allow you to turn on traversing")
	Progress := flag.Bool("progress", false, "This flag will allow you to turn on the progress bar")
	Discover := flag.Bool("discover", false, "Discover path/folder/file with already found traversal")
	Check := flag.Bool("check", false, "Check if a path/folder/file is internal")
	Test := flag.Bool("test", false, "Enable/Disable test mode only")
	flag.IntVar(&Depth, "depth", 5, "The crawl depth")

	flag.IntVar(&Threads, "c", 10, "The number of concurrent requests")
	flag.Parse()

	return true, *Crawl, *Silent, *Traverse, *Progress, *Test, *Discover, *Check
}
