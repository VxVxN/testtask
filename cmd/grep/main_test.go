package main

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintLines(t *testing.T) {
	re := regexp.MustCompile("line1")
	matchingLines := processLines([]string{"line1", "line2", "line3"}, re, 1, 0, false, false, false)
	assert.ElementsMatch(t, matchingLines, []string{"line1", "line2"})

	re = regexp.MustCompile("line3")
	matchingLines = processLines([]string{"line1", "line2", "line3"}, re, 0, 1, false, false, false)
	assert.ElementsMatch(t, matchingLines, []string{"line2", "line3"})

	re = regexp.MustCompile("line")
	matchingLines = processLines([]string{"line1", "Line2", "LINE3"}, re, 0, 0, false, true, false)
	assert.ElementsMatch(t, matchingLines, []string{"Line2", "LINE3"})

	re = regexp.MustCompile("line")
	matchingLines = processLines([]string{"line1", "Line2", "LINE3"}, re, 0, 0, true, true, false)
	assert.ElementsMatch(t, matchingLines, []string{"2"})

	re = regexp.MustCompile("line[23]")
	matchingLines = processLines([]string{"line1", "line2", "line3"}, re, 0, 0, false, false, true)
	assert.ElementsMatch(t, matchingLines, []string{"\x1b[32m2:\x1b[0mline2\n", "\x1b[32m3:\x1b[0mline3\n"})
}
