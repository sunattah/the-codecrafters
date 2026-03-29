package main

import (
	//"encoding/base64"
	"fmt"
	"strconv"
)

// func hexToDec() {

// return strconv.ParseBool(s)
// v, err := strconv.ParseBool(do)
// if err != nil {
// 	return true, nil
// } else {
// 	return v, nil
// }

// }
func hexToDec(s string) (int64, error) {
	// hex := []string{s,"FF", "1E"}
	file, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return 0, nil
	}
	return file, nil
}
func binToDec(s string) (int64, error) {
	file, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return 0, nil
	}
	return file, nil
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

// func decToBin(base int64, n int) string {
// 	d := strconv.FormatInt(base, n)
// 	return d

//}

func main() {
	var value string
	var operation string
	// var fish int64
	// var me int
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
		if value == "quit" {
			fmt.Println("Thanks and goodbye!")
			break
		}

		switch operation {

		case "1":
			result, err := hexToDec(value)
			if err != nil {
				fmt.Println("error converting hex", err)
				return
			}
			fmt.Printf("hex %s -> decimal:%d\n", value, result)

		case "2":
			result, err := binToDec(value)
			if err != nil {
				fmt.Println("error while converting binary", err)
				return
			}
			fmt.Printf("bin %s -> decimal:%d\n", value, result)
		case "3":
			hex, bin, err := decToHexAndBin(value)
			if err != nil {
				fmt.Println("error while converting binary", err)
				return
			}
			fmt.Printf("decimal: %s\n", value)
			fmt.Printf("binary: %s\n", hex)
			fmt.Printf("hexadecimal: %s\n", bin)

		default:
			fmt.Println("Invalid operation")
		}

	}

}
