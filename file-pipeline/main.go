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
)

func main() {
	var input string
	fmt.Println("<===THE RISE OF BANTITH IN CONGO===>")
	fmt.Println("press ENTER")
	fmt.Scan(&input)

	file, err := os.OpenFile("sample.txt", os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("invalid operation while opening file")
	}
	

	doe, err := os.ReadFile("sample.txt")
	 d := strings.Title(string(doe))
	 fmt.Println(d)
	if err != nil {
		fmt.Println("error while reading file")
	}
	fmt.Println(strings.ToUpper(string(doe)+"\n")) //+ (strings.ToLower(string(doe))))
	defer file.Close()
	file.WriteString("sample.txt")

}
