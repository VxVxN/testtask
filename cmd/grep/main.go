package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	after := flag.Int("A", 0, "Print N lines of trailing context after matching lines")
	before := flag.Int("B", 0, "Print N lines of leading context before matching lines")
	context := flag.Int("C", 0, "Print N lines of output context")
	count := flag.Bool("c", false, "Suppress normal output; instead print a count of matching lines for each input file")
	ignoreCase := flag.Bool("i", false, "Ignore case distinctions in patterns and input data")
	invert := flag.Bool("v", false, "Select non-matching lines")
	fixed := flag.Bool("F", false, "Interpret PATTERN as a fixed string, not a regular expression")
	numbered := flag.Bool("n", false, "Print line numbers with output lines")
	flag.Parse()

	if *context != 0 {
		after = context
		before = context
	}

	if len(os.Args) < 2 {
		log.Fatal("Usage: go-grep [options] pattern")
	}
	pattern := flag.Arg(0)

	lines, err := readLinesFromStdin()
	if err != nil {
		log.Fatalf("Failed to read from stdin: %v", err)
	}

	var re *regexp.Regexp
	if *fixed {
		re = regexp.MustCompile(regexp.QuoteMeta(pattern))
	} else {
		if *ignoreCase {
			re = regexp.MustCompile("(?i)" + pattern)
		} else {
			re = regexp.MustCompile(pattern)
		}
	}

	printLines(lines, re, *after, *before, *count, *invert, *numbered)
}

func readLinesFromStdin() ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func printLines(lines []string, re *regexp.Regexp, after, before int, count, invert, numbered bool) {
	var matchingLines []string
	for _, line := range lines {
		match := re.MatchString(line)
		if (invert && !match) || (!invert && match) {
			matchingLines = append(matchingLines, line)
		}
	}

	if count {
		fmt.Println(len(matchingLines))
		return
	}

	for i, line := range lines {
		match := re.MatchString(line)
		if (invert && !match) || (!invert && match) {
			printContext(lines, i, after, before, numbered)
		}
	}
}

func printContext(lines []string, index, after, before int, numbered bool) {
	start := index - before
	if start < 0 {
		start = 0
	}

	end := index + after + 1
	if end > len(lines) {
		end = len(lines)
	}

	for i := start; i < end; i++ {
		if numbered {
			fmt.Printf("\033[32m%d:\033[0m%s\n", index+1, lines[i])
			continue
		}
		fmt.Println(lines[i])
	}
}
