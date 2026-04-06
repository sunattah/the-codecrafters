package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func capWords(s string) string {
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		words[i] = strings.Title(strings.ToLower(words[i]))
	}
	return strings.Join(words, " ")
}

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func todo(s string) string {
	return strings.ReplaceAll(s, "TODO", "ACTION")
}

func class(s string) string {
	return strings.ReplaceAll(s, "CLASSIFIED", "[REDACTED]")
}

func replace(s string) string {
	s = strings.ReplaceAll(s, " -", "")
	s = strings.ReplaceAll(s, " ", " ")
	return s
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
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
		result := line

		result = todo(result)
		result = class(result)
		result = replace(result)
		result = reverse(result)
		result = capWords(result)

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
