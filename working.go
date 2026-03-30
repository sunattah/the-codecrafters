package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	num1, num2 float64 // integer storer
	p_choice   int     // proceed choice whether to do another calculation or not
)

func main() {

	reader := bufio.NewReader(os.Stdin)
start:
	fmt.Printf("CHOOSE AN OPERATION TO EXECUTE\n 1. Addition\n 2. Subtraction\n 3. Multiplication\n 4. Division\n 5. Exit\n 6. More info\n")
	fmt.Print("Select an operation to continue: ")
	reader, _ := reader.ReadString('\n')
	reader = strings.TrimSpace(reader)
	op, _ := strconv.Atoi(reader)

	for {

		switch op {
		case 1:

		case1FirstNumber:
			fmt.Print("Enter first number: ")
			_, err := fmt.Scanln(&num1)
			if err != nil {
				fmt.Println("Enter a number and not alphabet")
				goto case1FirstNumber
			}
		case1SecondNumber:
			fmt.Print("Enter second number: ")
			_, err = fmt.Scanln(&num2)
			if err != nil {
				fmt.Println("Enter a number and not alphabet")
				goto case1SecondNumber
			}
			result := float64(num1) + float64(num2)

			fmt.Println(result)
			fmt.Println()

			fmt.Print("Do you want add another set of numbers?\n 1. Yes, I want to.\n 2. No, Quit.\n 3. Go to Menu\n")
			fmt.Print("Pick your choice: ")
			fmt.Scanln(&p_choice)

			if p_choice == 1 {
				continue
			} else if p_choice == 2 {
				return
			} else if p_choice == 3 {
				goto start
			} else {
				fmt.Println("Choose from the options above")
			}

		case 2:
		case2FirstNumber:
			fmt.Print("Enter first number: ")
			_, err := fmt.Scanln(&num1)
			if err != nil {
				fmt.Println("Enter a number and not alphabet")
				goto case2FirstNumber
			}
		case2SecondNumber:
			fmt.Print("Enter second number: ")
			_, err = fmt.Scanln(&num2)
			if err != nil {
				fmt.Println("Enter a number and not alphabet")
				goto case2SecondNumber
			}
			result := float64(num1) - float64(num2)
			fmt.Println(result)
			//fmt.Printf("%f - %f is = %f\n", num1, num2, result)

			fmt.Print("Do you want subtract another set of numbers?\n 1. Yes, I want to.\n 2. No, Quit.\n 3. Go to Menu\n")
			fmt.Print("Pick your choice: ")
			fmt.Scanln(&p_choice)

			if p_choice == 1 {
				continue
			} else if p_choice == 2 {
				return
			} else if p_choice == 3 {
				goto start
			} else {
				fmt.Println("Choose from the options above")
			}

		case 3:
		case3FirstNumber:
			fmt.Print("Enter first number: ")
			_, err := fmt.Scanln(&num1)
			if err != nil {
				fmt.Println("Enter a number and not alphabet")
				goto case3FirstNumber
			}
		case3SecondNumber:
			fmt.Print("Enter second number: ")
			_, err = fmt.Scanln(&num2)
			if err != nil {
				fmt.Println("Enter a number and not alphabet")
				goto case3SecondNumber
			}
			result := float64(num1) * float64(num2)
			//fmt.Printf("%d x %d is = %d\n", num1, num2, result)
			fmt.Println(result)

			fmt.Print("Do you want multiply another set of numbers?\n 1. Yes, I want to.\n 2. No, Quit.\n 3. Go to Menu\n")
			fmt.Print("Pick your choice: ")
			fmt.Scanln(&p_choice)

			if p_choice == 1 {
				continue
			} else if p_choice == 2 {
				return
			} else if p_choice == 3 {
				goto start
			} else {
				fmt.Println("Choose from the options above")
			}

		case 4:
		case4FirstNumber:
			fmt.Print("Enter first number: ")
			_, err := fmt.Scanln(&num1)
			if err != nil {
				fmt.Println("Enter a number and not alphabet")
				goto case4FirstNumber
			}
		case4SecondNumber:
			fmt.Print("Enter second number: ")
			_, err = fmt.Scanln(&num2)
			if err != nil {
				fmt.Println("Enter a number and not alphabet")
				goto case4SecondNumber
			}
			if num2 == 0 {
				fmt.Println("Divisor can't be zero")
				goto case4SecondNumber
			}
			result := float64(num1) / float64(num2)
			//fmt.Printf("%d ÷ %d is = %d\n", num1, num2, result)
			fmt.Println(result)

			fmt.Print("Do you want divide another set of numbers?\n 1. Yes, I want to.\n 2. No, Quit.\n 3. Go to Menu\n")
			fmt.Print("Pick your choice: ")
			fmt.Scanln(&p_choice)

			if p_choice == 1 {
				continue
			} else if p_choice == 2 {
				return
			} else if p_choice == 3 {
				goto start
			} else {
				fmt.Println("Choose from the options above")
			}

		case 5:
			fmt.Println("Exiting...")
			return

		case 6:
			fmt.Println()
			fmt.Println(" Addition: add your input together\n Subtraction: Minus a number from another\n Multiplication: Multiply numbers to get result\n Division: Divides a number by the other\n ")
			return

		default:
			fmt.Println("Enter a valid operation: ")
			goto start

		}

	}
}



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
func hexToDec(s string) (string) {
	// hex := []string{s,"FF", "1E"}
	_, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		return "error: please input the appropriate value"
	} else {
		return "good"
	}
}

func binToDec(s string) (string) {
	_, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		return "error: please insert the appropriate value"
	}
	return "good"
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
		if operation == "quit" {
			fmt.Println("Thanks and goodbye!")
			break
		}

		switch operation {

		case "1":
			fmt.Println(binToDec(""))
			
				fmt.Printf("hex %s -> decimal:%s\n", value)
			

		case "2":
			fmt.Println(hexToDec(""))

			fmt.Printf("bin %s -> decimal:%s\n", value)
		case "3":
			hex, bin, err := decToHexAndBin(value)
			if err != nil {
				fmt.Println("error while converting binary")

			}

			fmt.Printf("decimal: %s\n", value)
			fmt.Printf("binary: %s\n", hex)
			fmt.Printf("hexadecimal: %s\n", bin)
		default:
			fmt.Println("Invalid operation")

		}

	}

}
