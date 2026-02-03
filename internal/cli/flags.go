package cli

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	Keyword         string
	CaseInsensitive bool
	CountOnly       bool
	TopN            int
	Machine         bool
	Help            bool
}

// ParseFlags parses command-line flags and returns a Config.
// It registers both short and long forms for the main flags.
func ParseFlags() (*Config, error) {
	var (
		keyword         string
		caseInsensitive bool
		countOnly       bool
		topN            int
		machine         bool
		help            bool
	)

	// Keyword (short and long)
	flag.StringVar(&keyword, "k", "", "Keyword to search for (substring match)")
	flag.StringVar(&keyword, "keyword", "", "Keyword to search for (substring match)")

	// Case-insensitive (short and long)
	flag.BoolVar(&caseInsensitive, "i", false, "Case-insensitive matching")
	flag.BoolVar(&caseInsensitive, "insensitive", false, "Case-insensitive matching")

	// Count only
	flag.BoolVar(&countOnly, "count", false, "Only print the count of matching lines")

	// Top N frequent lines
	flag.IntVar(&topN, "top", 0, "Show top N most frequent matching lines (0 disables)")

	// Machine-readable output
	flag.BoolVar(&machine, "machine", false, "Machine-readable output (tab-separated for top)")

	// Help
	flag.BoolVar(&help, "help", false, "Show this help")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] [FILE]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Read from FILE or stdin. Filter lines containing keyword.\n\n")
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	config := &Config{
		Keyword:         keyword,
		CaseInsensitive: caseInsensitive,
		CountOnly:       countOnly,
		TopN:            topN,
		Machine:         machine,
		Help:            help,
	}
	return config, nil
}// feb iteration 29
// feb iteration 30
// feb iteration 31
