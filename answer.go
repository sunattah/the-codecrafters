ackage main 

import(
	  "fmt"
)


func main() {
	var number1 float64 
	var number2 float64 
	var operator string 

    fmt.Println()

	fmt.Println("enter value:")
	fmt.Scan(&number1)


	fmt.Println("enter value:")
	fmt.Scan(&number2)

	fmt.Println("choose an operator: + - * /")
	fmt.Scan(&operator)

	switch operator {
	case "+":
	fmt.Printf("%f %s %f = %f", number1, operator, number2,  number1 + number2 )
	case "-":
	fmt.Printf("%f %s %f = %f", number1, operator, number2,  number1 - number2 )
	case "*":
	fmt.Printf("%f %s %f = %f", number1, operator, number2,  number1 + number2 )
	case "/":
	fmt.Printf("%f %s %f = %f", number1, operator, number2,  number1 / number2 )
	if number2 == 0.0 {
		fmt.Println("divide by zero")
	}
       default:
	fmt.Println("invalid operator")
	}


}
