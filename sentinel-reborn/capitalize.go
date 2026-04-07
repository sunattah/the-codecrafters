package main

import (
	"strconv"
	"strings"
)

func capitalizeWord(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(word[:1]) + word[1:]
}

func getCapCount(token string) int {
	
	if strings.EqualFold(token, "(cap)") {
		return 1
	}

	token = strings.ToLower(token)
	token = strings.TrimPrefix(token, "(cap,")
	token = strings.TrimSuffix(token, ")")
	token = strings.TrimSpace(token)

	n, err := strconv.Atoi(token)
	if err != nil || n <= 0 {
		return 1
	}
	return n
}

func capitalize(s string) string {
	words := strings.Fields(s)
	result := []string{}

	for _, w := range words {
		if strings.HasPrefix(strings.ToLower(w), "(cap") {
			count := getCapCount(w)
			for i := 0; i < count && i < len(result); i++ {
				idx := len(result) - 1 - i
				result[idx] = capitalizeWord(result[idx])
			}
		} else {
			result = append(result, w)
		}
	}

	return strings.Join(result, " ")
}