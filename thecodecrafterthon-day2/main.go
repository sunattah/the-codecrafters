package main

import (
	//"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

var value string
var operation string

// func hexToDec() {

// return strconv.ParseBool(s)
// v, err := strconv.ParseBool(do)
// if err != nil {
// 	return true, nil
// } else {
// 	return v, nil
// }

// }
func hexToDec(s string) {

	file, err := strconv.ParseInt(s, 16, 64)
	if err == nil {
		fmt.Printf("hex %s -> decimal:%d\n", value, file)
	} else {
		fmt.Println("An error occured! ")
		fmt.Println(err)
	}
}
func binToDec(s string) {
	file, err := strconv.ParseInt(s, 2, 64)
	if err == nil {
		fmt.Printf("bin %s -> decimal:%d\n", value, file)
	} else {
		fmt.Println("an error occured")
		fmt.Println(err)
	}

}
func decToHexAndBin(s string) (string, string, error) {
	file, err := strconv.ParseInt(s, 10, 64)
	if err != nil {

		return "", "", nil
	}
	y := strconv.FormatInt(file, 2)
	d := strconv.FormatInt(file, 16)
	return y, d, nil

}

func main() {

	for {
		fmt.Println("Enter the value: ")

		fmt.Scan(&value)
		if value == "quit" {
			fmt.Println("Thanks and goodbye!")
			break
		}

		fmt.Println("Enter the operation: ")
		fmt.Println("1.convert from hex to Decimal")
		fmt.Println("2.convert from bin to decimal")
		fmt.Println("3.convert from dec to binary and hexadecimal")
		fmt.Scan(&operation)
		if operation == "quit" {
			fmt.Println("Thanks and goodbye!")
			break
		}

		switch operation {

		case "1":
			hexToDec(value)

		case "2":
			binToDec(value)

		case "3":
			hex, bin, err := decToHexAndBin(value)
			if err != nil {
				fmt.Println("error while converting binary", err)
				return
			}
			fmt.Printf("decimal: %s\n", value)
			fmt.Printf("binary: %s\n", strings.ToUpper(hex))
			fmt.Printf("hexadecimal: %s\n", strings.ToUpper(bin))

		default:
			fmt.Println("Invalid operation")
		}

	}

}
