package main

import (
	"fmt"
	"strings"
)

func upper(w string) string {
	s := strings.Fields(w)
	for i := 0; i < len(s); i++ {
		s[i] = strings.ToUpper(s[i][0:])
	}
	return strings.Join(s, " ")
}
func cap(a string) string {
	d := strings.Fields(a)
	for i := range d {
		d[i] = strings.Title(string(d[i][0:]))

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


func main() {
	var input string
	for {
		fmt.Println("<== Choose from the available inputs ==>")
		fmt.Println("upper")
		fmt.Println("cap")
		fmt.Println("low")
		fmt.Println("reverse")
		fmt.Scan(&input)
		if input == "exit" {
			fmt.Println("Shutting down String Transformer. Goodbye.")
			break
		}

		switch input {
		case "upper":
			upper("")
		case "cap":
			cap("")
		case "low":
			low("")
		case "reverse":
			reverse("")

		default:
			fmt.Println("An invalid input: please type the appropriate value")

		}

	}
}
