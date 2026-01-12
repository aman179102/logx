package core

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Processor struct {
	Keyword         string
	CaseInsensitive bool
	CountOnly       bool
	TopN            int
	Machine         bool
	Out             io.Writer // where to write output (usually os.Stdout)
}

type Result struct {
	MatchCount int // total number of matching lines
}

// Process runs the appropriate operation based on processor flags.
func (p *Processor) Process(scanner *bufio.Scanner) (*Result, error) {
	if p.TopN > 0 {
		return p.processTop(scanner)
	}
	if p.CountOnly {
		return p.processCount(scanner)
	}
	return p.processFilter(scanner)
}

// processFilter prints matching lines.
func (p *Processor) processFilter(scanner *bufio.Scanner) (*Result, error) {
	var matchCount int
	compare := p.makeCompareFunc()
	for scanner.Scan() {
		line := scanner.Text()
		if compare(line) {
			matchCount++
			fmt.Fprintln(p.Out, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &Result{MatchCount: matchCount}, nil
}

// processCount prints only the count of matching lines.
func (p *Processor) processCount(scanner *bufio.Scanner) (*Result, error) {
	var matchCount int
	compare := p.makeCompareFunc()
	for scanner.Scan() {
		line := scanner.Text()
		if compare(line) {
			matchCount++
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if p.Machine {
		fmt.Fprintf(p.Out, "%d\n", matchCount)
	} else {
		fmt.Fprintf(p.Out, "%d\n", matchCount)
	}
	return &Result{MatchCount: matchCount}, nil
}

// processTop prints the top N most frequent matching lines.
func (p *Processor) processTop(scanner *bufio.Scanner) (*Result, error) {
	freq := make(map[string]int)
	compare := p.makeCompareFunc()

	for scanner.Scan() {
		line := scanner.Text()
		if compare(line) {
			freq[line]++
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Calculate total matches (sum of frequencies)
	totalMatches := 0
	for _, cnt := range freq {
		totalMatches += cnt
	}

	// Build slice for sorting
	type pair struct {
		Line  string
		Count int
	}
	pairs := make([]pair, 0, len(freq))
	for line, cnt := range freq {
		pairs = append(pairs, pair{Line: line, Count: cnt})
	}

	// Sort: descending count, then ascending line for determinism
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].Count == pairs[j].Count {
			return pairs[i].Line < pairs[j].Line
		}
		return pairs[i].Count > pairs[j].Count
	})

	// Take top N
	n := p.TopN
	if n > len(pairs) {
		n = len(pairs)
	}
	topPairs := pairs[:n]

	// Output
	if p.Machine {
		for _, pair := range topPairs {
			fmt.Fprintf(p.Out, "%d\t%s\n", pair.Count, pair.Line)
		}
	} else {
		for _, pair := range topPairs {
			fmt.Fprintf(p.Out, "%d %s\n", pair.Count, pair.Line)
		}
	}

	return &Result{MatchCount: totalMatches}, nil
}

// makeCompareFunc returns a function that checks if a line matches the keyword.
func (p *Processor) makeCompareFunc() func(string) bool {
	if p.Keyword == "" {
		// No keyword: match everything
		return func(string) bool { return true }
	}
	if p.CaseInsensitive {
		kw := strings.ToLower(p.Keyword)
		return func(line string) bool {
			return strings.Contains(strings.ToLower(line), kw)
		}
	}
	return func(line string) bool {
		return strings.Contains(line, p.Keyword)
	}
}// jan iteration 1
// jan iteration 2
// jan iteration 3
// jan iteration 4
// jan iteration 5
// jan iteration 6
// jan iteration 7
// jan iteration 8
// jan iteration 9
