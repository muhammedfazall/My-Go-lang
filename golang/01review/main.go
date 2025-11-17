package main

import "fmt"

func main() {
	num1, num2 := 0, 0
	op := ""

	fmt.Println("Enter first operator")
	fmt.Scan(&num1)

	fmt.Println("Enter Operation (+,-,*,/)")
	fmt.Scan(&op)

	fmt.Println("Enter second operator")
	fmt.Scan(&num2)

	result := 0

	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error! division by zero")
			return
		}
		result = num1 / num2
	}

	fmt.Println("Result:", result)

}
