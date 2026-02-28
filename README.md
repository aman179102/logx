# logx — a small Unix-style log analyzer

`logx` is a minimal command-line tool for filtering, counting, and summarizing log files.  
It reads from a file or standard input, matches lines containing a keyword, and can report
simple frequency statistics. The goal is to stay predictable, script-friendly, and boring.

## Why logx?

In many day-to-day workflows, I only need quick answers like:

- Does this log contain errors?
- How many times did this message appear?
- What is the most common log line?

These are all possible with `grep` and `awk`, but they usually require writing small scripts
each time. `logx` exists as a deliberately simple alternative for these common cases.

It focuses on:
- Substring matching (not regex)
- Clear exit codes for scripting
- Built-in frequency analysis
- Zero configuration and no dependencies

## Design philosophy

- **Do one thing well**  
  Filter and summarize log lines. Nothing more.

- **Streaming by default**  
  Files are processed line by line without loading everything into memory.

- **Predictable behavior**  
  Deterministic output ordering and well-defined exit codes.

- **Composable**  
  Works naturally with pipes and shell redirections.

- **Boring code**  
  Written in plain Go using only the standard library, easy to read and audit.

## Installation

### Build from source

```bash
git clone https://github.com/aman179102/logx
cd logx
go build -o logx ./cmd/logx
sudo mv logx /usr/local/bin/
// polish 48
// polish 49
// polish 50
// polish 51
// polish 52
// polish 53
// polish 54
// polish 55
// polish 56
