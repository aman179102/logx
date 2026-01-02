package main

import (
	"flag"
	"fmt"
	"os"

	"logx/internal/cli"
	"logx/internal/core"
	"logx/internal/io"
)

func main() {
	config, err := cli.ParseFlags()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing flags: %v\n", err)
		os.Exit(2)
	}
	if config.Help {
		flag.Usage()
		os.Exit(0)
	}

	// Determine input source
	args := flag.Args()
	var path string
	if len(args) > 0 {
		path = args[0]
		if len(args) > 1 {
			fmt.Fprintf(os.Stderr, "Error: too many arguments\n")
			os.Exit(2)
		}
	}

	// Open input (streaming)
	scanner, closer, err := io.NewLineScanner(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening input: %v\n", err)
		os.Exit(2)
	}
	defer closer()

	// Create processor
	proc := &core.Processor{
		Keyword:         config.Keyword,
		CaseInsensitive: config.CaseInsensitive,
		CountOnly:       config.CountOnly,
		TopN:            config.TopN,
		Machine:         config.Machine,
		Out:             os.Stdout,
	}

	// Process
	result, err := proc.Process(scanner)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error processing: %v\n", err)
		os.Exit(2)
	}

	// Exit with appropriate code
	if result.MatchCount > 0 {
		os.Exit(0)
	} else {
		os.Exit(1)
	}
}