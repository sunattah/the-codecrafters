// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: [Your Name]
// Squad:  [Your Squad Name]

// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: [Squad Name]
// ───────────────────────────────────────────
// Input line types:
//   [list what your squad agreed on]
//
// Transformation rules (in order):
//   1. [Rule 1]
//   2. [Rule 2]
//   3. [Rule 3]
//   4. [Rule 4]
//   5. [Rule 5]
//
// Output format:
//   [Header: yes/no — exact text if yes]
//   [Line numbering format]
//   [Summary block: yes/no — fields if yes]
//
// Terminal summary fields:
//   [List what your squad agreed on]
// ═══════════════════════════════════════════

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	// "strings"
)

func cap(s string) string {
	d := strings.Fields(s)
	for i := 0; i < len(d); i++ {
		d[i] = strings.Title(string(strings.ToLower(d[i][0:])))
	}
	return strings.Join(d, " ")
}
func toUpper(s string) string {
	return strings.ToUpper(s)
}
func todo(s string) string {
	t := "TODO"
	d := strings.ReplaceAll(s, t, "ACTION")
	return d

}
func class(s string) string {
	t := " CLASSIFIED"
	d := strings.ReplaceAll(s, t, " [REDACTED]")
	return d
}
func replace(s string) string {
	t := strings.ReplaceAll(s, " -", "")
	a := strings.ReplaceAll(t, " ", "")
	return a
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go input.txt output.txt")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	data, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	var processed []string

	for _, line := range lines {
		result := cap(line)
		result = toUpper(result)
		result = todo(result)
		result = class(result)
		processed = append(processed, result)
	}

	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer out.Close()

	writer := bufio.NewWriter(out)

	for i, line := range processed {
		fmt.Fprintf(writer, "%d. %s\n", i+1, line)
	}

	writer.Flush()

	fmt.Println("Done! Check", outputFile, "for your results.")
}
