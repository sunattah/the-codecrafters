package main

import (
	//"encoding/base64"
	//iuy65r4e3"errors"
	"fmt"
	"strconv"
	"strings"
)

var value string
var operation string

func hexToDec() {
	fmt.Println("Enter the value: ")
	fmt.Scan(&value)
	file, err := strconv.ParseInt(value, 16, 64)
	if err == nil {
		fmt.Printf("hex %s -> decimal:%d\n", value, file)
	} else {
		fmt.Println("An error occured! ")
		fmt.Println(err)
	}
}
func binToDec() {
	fmt.Println("Enter the value: ")
	fmt.Scan(&value)
	file, err := strconv.ParseInt(value, 2, 64)
	if err == nil {
		fmt.Printf("hex %s -> decimal:%d\n", value, file)
	} else {
		fmt.Println("An error occured! ")
		fmt.Println(err)
	}

}
func decToHexAndBin() {
	fmt.Println("Enter the value: ")
	fmt.Scan(&value)
	new, err := strconv.Atoi(value)
	if err == nil {
		file1 := strconv.FormatInt(int64(new), 2)
		file2 := strconv.FormatInt(int64(new), 16)
		fmt.Printf("binary: %s\n", file1)
		fmt.Printf("hexdecimal: %s\n", strings.ToUpper(file2))

	} else {
		fmt.Println("an error occured: try again")

	}
	// if err != nil {

	// 	return "", "", nil
	// }
	// y := strconv.FormatInt(file, 2)
	// d := strconv.FormatInt(file, 16)
	// return y, d, nil

}

func main() {

	for {
		fmt.Println("Enter the operation: ")
		fmt.Println("1.convert from hex to Decimal")
		fmt.Println("2.convert from bin to decimal")
		fmt.Println("3.convert from dec to binary and hexadecimal")
		fmt.Scanln(&operation)
		if operation == "quit" {
			fmt.Println("Thanks and goodbye!")
			break
		}

		switch operation {

		case "1":
			hexToDec()

		case "2":
			binToDec()

		case "3":
			decToHexAndBin()

		default:
			fmt.Println("Invalid operation")
		}

	}

}
