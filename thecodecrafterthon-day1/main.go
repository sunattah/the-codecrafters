package main

import (
	"fmt"
)

func main() {
	var number1 float64
	var number2 float64
	var operator string
	//var gos string
	//var race string
	//var dog string
	for {

		fmt.Println("enter first value:")
		fmt.Scan(&number1)

		fmt.Println("enter second value:")
		fmt.Scan(&number2)

		fmt.Println("choose an operator: + - * /, help")
		fmt.Scan(&operator)

		if operator == "quit" {
			fmt.Printf("goodbye and thankyou:")
			break
		}
		// switch gos {
		// case"[a-z]":

		// 	fmt.Println("invalid")
		//    if gos == "" {
		// 	fmt.Println("valid")
		//    }

		// }

		switch operator {

		case "a,b,c,d,e,f,g,h,i,j,k,l,m,n,o,p,q,r,s,t,u,v,w,x,y,z":
			fmt.Println("not numbers")
		case "help":
			fmt.Println("go to https//go package.com for help")
		case "+":
			fmt.Printf("%.3f %s %.3f = %.3f\n", number1, operator, number2, number1+number2)
		case "-":
			fmt.Printf("%.3f %s %.3f = %.3f\n", number1, operator, number2, number1-number2)
		case "*":
			fmt.Printf("%.3f %s %.3f = %.3f\n", number1, operator, number2, number1+number2)
		case "/":
			fmt.Printf("%.3f %s %.3f = %.3f\n", number1, operator, number2, number1/number2)
			if number2 == 0.00 {
				fmt.Println("not divisible by zero")
			}
		// case "quit":
		// fmt.Printf("good bye, thankyou")

		default:
			fmt.Println("invalid operator")

		}
		continue

	}
}
