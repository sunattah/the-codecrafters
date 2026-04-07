  package main

import "fmt"

func main() {
	var input1 int
	var input2 int
	var operator string
	for {
		fmt.Println("\033[34mSCIENTIFIC CAL READY TO SOLVE:\033[0m ")
		fmt.Println("\033[33m<== please insert any value ==>\033[0m")

		fmt.Print("input the first value: ")
		fmt.Scan(&input1)

		fmt.Println("input the second value: ")

		fmt.Scan(&input2)

		fmt.Println("choose an operator to use: + / * - help ")
		fmt.Scan(&operator)
		if operator == "quit" {
			fmt.Println("\033[31mThank you for using this calc and goodbye\033[0m")
			break
		}

		switch operator {
		case "help":
			fmt.Println("check the calc documentation")
		case "+":
			fmt.Printf("%d %s %d = %d\n", input1, operator, input2, input1+input2)
		case "-":
			fmt.Printf("%d %s %d = %d\n", input1, operator, input2, input1-input2)
		case "*":
			fmt.Printf("%d %s %d = %d\n", input1, operator, input2, input1*input2)
		case "/":

			if input2 == 0 {
				fmt.Println("it is not divisible")
			} else {
				fmt.Printf("%d %s %d = %d\n", input1, operator, input2, input1/input2)

			}

		default:
			fmt.Println("error, it is not the proper operation, please insert the correct operation, and value to prevent panicing")

		}
		continue

	}
}
