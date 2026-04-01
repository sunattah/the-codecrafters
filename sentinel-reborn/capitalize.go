package main

import (
	"fmt"
	"strings"
	
)

func capitalize(s string) string {
	e := "(cap)"

	d := strings.Fields(s)

	for i := 0; i < len(d); i++ {
		if d[i] == e && i > 0 {
			d[i-1] = strings.Title(d[i-1])
			d = append(d[:i], d[i+1:]...)
		}
	}
	return strings.Join(d, " ")

}

//	func de(s string) []string {
//		d := strings.Fields(s)
//		for i := range d {
//			if i < len(d)-1 {
//				d[i] = strings.Title(d[i])
//			}
//			return d
//		}
//		return d
//	}
func main() {
	fmt.Println(capitalize("hgfe bhdfgduo fhdf (cap 6)"))
	// fmt.Println(de("this"),2)
}
