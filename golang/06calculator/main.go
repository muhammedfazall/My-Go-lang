package main

import "fmt"

func main() {
	var num1, num2 float64
	var op string

	fmt.Println("Simple CLI calculator(Go)")

	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter Operator (+,-,*,/) ")
	fmt.Scanln(&op)

	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	result := 0.0

	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0{
			fmt.Println("ERROR!: Division by zero")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid Operation: Use + , - , * or /")
	}

	fmt.Printf("Result is : %v", result)
}
