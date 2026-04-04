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

	file, err := os.OpenFile("sample.txt", os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("invalid operation while opening file")
	}
	defer file.Close()
	doe, err := os.ReadFile("sample.txt")
	if err != nil {
		fmt.Println("error while reading file")
	}
	fmt.Println(string(doe)) 
	file.WriteString("sample.txt")

}
