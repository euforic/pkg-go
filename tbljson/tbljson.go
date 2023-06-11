package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// TableToMapSlice converts an ASCII table to a slice of maps.
func TableToMapSlice(table string) ([]map[string]string, error) {
	scanner := bufio.NewScanner(strings.NewReader(table))
	var header []string
	var rows [][]string
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 0 {
			continue
		}
		if len(header) == 0 {
			header = fields
			continue
		}
		rows = append(rows, fields)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	result := make([]map[string]string, len(rows))
	for i, row := range rows {
		rowMap := make(map[string]string)
		for j, value := range row {
			if j < len(header) {
				rowMap[header[j]] = value
			} else {
				rowMap[fmt.Sprintf("Column %d", j+1)] = value
			}
		}
		result[i] = rowMap
	}
	return result, nil
}

func main() {
	// Read from standard input
	scanner := bufio.NewScanner(os.Stdin)
	var output string
	for scanner.Scan() {
		output += scanner.Text() + "\n"
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	// Convert the output to a map slice
	mapSlice, err := TableToMapSlice(output)
	if err != nil {
		panic(err)
	}

	// Print the JSON output
	jsonOutput, err := json.MarshalIndent(mapSlice, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonOutput))
}
