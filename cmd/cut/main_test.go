package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFields(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{"1,2,3", []int{0, 1, 2}},
		{"2,4,6", []int{1, 3, 5}},
		{"1,1,2", []int{0, 0, 1}}, // check duplicate
	}

	for _, test := range tests {
		result := parseFieldIndexes(test.input)
		assert.Equal(t, test.expected, result, "Unexpected result for input: %s", test.input)
	}
}

func TestExtractFields(t *testing.T) {
	tests := []struct {
		columns  []string
		indexes  []int
		expected string
	}{
		{[]string{"a", "b", "c"}, []int{0, 1}, "a\tb"},
		{[]string{"a", "b", "c"}, []int{1, 2}, "b\tc"},
		{[]string{"a", "b", "c"}, []int{2}, "c"},
		{[]string{"a", "b", "c"}, []int{3}, ""}, // index out of range
	}

	for _, test := range tests {
		result := extractFields(test.columns, test.indexes)
		assert.Equal(t, test.expected, result, "Unexpected result for columns: %v, indexes: %v", test.columns, test.indexes)
	}
}
