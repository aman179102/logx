# logx – a simple log file filter and analyzer

`logx` is a Unix-style command-line tool for filtering, counting, and analyzing log files. It reads from a file or standard input, matches lines containing a keyword, and can show the most frequent matching lines. Designed to be simple, predictable, and suitable for scripting.

## Why logx?

While `grep` is excellent for pattern matching, it doesn't offer built-in frequency analysis or script-friendly exit codes. `awk` can do frequency counts but requires writing non-trivial scripts. `logx` fills a small niche:

- Filter lines by a keyword (substring match).
- Count matches and exit with 0/1 to indicate presence/absence.
- Show the top N most frequent matching lines.
- Machine-readable output for easy integration with other tools.
- Works offline, no dependencies, handles huge files.

## Design Philosophy

- **Do one thing well**: Filter and analyze log lines.
- **Streaming**: Never load the whole file into memory.
- **Predictable**: Deterministic sorting, clear exit codes.
- **Composable**: Works with pipes and redirections.
- **Boring code**: Written in pure Go, easy to audit and maintain.

## Installation

Download the static binary for Linux from the [releases page](https://example.com) (no hosting provided – you build it yourself or use your package manager).  

To build from source:
```bash
git clone https://github.com/yourorg/logx
cd logx
go build -o logx ./cmd/logx
sudo mv logx /usr/local/bin