// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: [sunday attah]
// Squad:  [pointers]

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func upper(w string) string {
	return strings.ToUpper(w)

}
func cap(a string) string {
	d := strings.Fields(a)
	for i := range d {
		d[i] = strings.Title(string(strings.ToLower(d[i][0:])))

	}
	return strings.Join(d, " ")
}
func low(d string) string {
	s := strings.Fields(d)
	for i := range s {
		s[i] = strings.ToLower(s[i][:])
	}
	return strings.Join(s, " ")
}
func reverse(l string) string {
	f := strings.Split(l, "")
	for i, j := 0, len(f)-1; i < j; i, j = i+1, j-1 {
		f[i], f[j] = f[j], f[i]
	}
	return strings.Join(f, "")
}
func underscore(s string) string {
	f := strings.ReplaceAll(s, "! ", " ")

	d := strings.Split(strings.ToLower(f), " ")
	return strings.Join(d, "_")
}

func title() {
	var smallWords = []string{"a", "an", "the", "and", "but",
		"or", "for", "nor", "on", "at", "to", "by",
		"in", "of", "up", "as", "is", "it"}

	text := bufio.NewScanner(os.Stdin)
	fmt.Println("sorry input another text/thesame text")
	if text.Scan() {
		line := text.Text()
		words := strings.Fields(line)

		for i, word := range words {
			read := slices.Contains(smallWords, strings.ToLower(word))
			if i == 0 || !read {
				words[i] = strings.ToUpper(word[:1]) + strings.ToLower(word[1:])
			} else {
				words[i] = strings.ToLower(word)
			}
		}
		fmt.Println(strings.Join(words, " "))
	}
}

func main() {
	var input string

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\033[33m<== Choose from the availables below ==>\033[0m.")
		fmt.Println("upper")
		fmt.Println("cap")
		fmt.Println("low")
		fmt.Println("reverse")
		fmt.Println("underscore")
		fmt.Println("title")
		fmt.Println("== you must choose from the above options ==")
		fmt.Scan(&input)
		if input == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye.")
			break
		}

		fmt.Print("\033[34m input your value: \033[0m")

		scanner.Scan()
		value := scanner.Text()
		value = strings.TrimSpace(value)

		switch input {
		case "upper":
			fmt.Println(upper(value))
		case "cap":
			fmt.Println(cap(value))
		case "low":
			fmt.Println(low(value))
		case "reverse":
			fmt.Println(reverse(value))
		case "underscore":
			fmt.Println(underscore(value))
		case "title":
			title()

		default:
			fmt.Println("An invalid input: please type the appropriate value")

		}

	}
}
