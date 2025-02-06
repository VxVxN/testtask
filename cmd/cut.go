package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var fields string
	var delimiter string
	var separated bool

	flag.StringVar(&fields, "f", "", "Select the fields")
	flag.StringVar(&delimiter, "d", "\t", "Use a different separator")
	flag.BoolVar(&separated, "s", false, "Delimited lines only")
	flag.Parse()

	if fields == "" {
		log.Fatal("No fields selected")
	}

	fieldIndexes := parseFields(fields)
	if delimiter == "" {
		delimiter = "\t"
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if separated && !strings.Contains(line, delimiter) {
			continue
		}

		columns := strings.Split(line, delimiter)
		output := extractFields(columns, fieldIndexes)

		if output != "" {
			fmt.Println(output)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning input: %s", err)
	}
}

func parseFields(rawFields string) []int {
	fields := strings.Split(rawFields, ",")

	fieldIndexes := make([]int, 0, len(fields))
	for _, field := range fields {
		var index int
		_, err := fmt.Sscanf(field, "%d", &index)
		if err == nil && index > 0 {
			fieldIndexes = append(fieldIndexes, index-1)
		} else {
			log.Fatalf("Error parsing field: %s", field)
		}
	}
	return fieldIndexes
}

func extractFields(columns []string, fieldIndexes []int) string {
	var output []string
	for _, index := range fieldIndexes {
		if index < len(columns) {
			output = append(output, columns[index])
		}
	}
	return strings.Join(output, "\t")
}
