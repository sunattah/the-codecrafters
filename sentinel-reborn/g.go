package main

import (
	"fmt"
	"strconv"
	"strings"
)

func capitalize(s string) string {
	d := strings.Fields(s)

	for i := 0; i < len(d); i++ {
		if d[i] == "(cap" && i+1 < len(d) {
			numStr := strings.TrimSuffix(d[i+1], ")")
			n, err := strconv.Atoi(numStr)
			if err != nil {
				continue
			}
			start := i - n
			if start < 0 {
				start = 0
			}
			for in := start; in < i; in++ {
				d[in] = strings.Title(d[in])
			}
			d = append(d[:i], d[i+2:]...)
			i--
		}
	}
	return strings.Join(d, " ")
}

func main() {
	fmt.Println(capitalize("hgfe bhdfgduo fhdf (cap 6)"))
}
