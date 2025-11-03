package main

import "fmt"

func divide(n1 int,n2 int) int {
	if n2 == 0 {
		panic("Cant divide by zero")
	}
	return n1/n2
}

func main()  {
	defer func ()  {
		if r := recover(); r != nil{
			fmt.Println("Panic Recovered:",r)
		}	
	}()

	fmt.Println(divide(50,0))
}

	// op := ""

	// fmt.Println("Enter first operator")
	// fmt.Scan(&num1)

	// fmt.Println("Enter Operation (+,-,*,/)")
	// fmt.Scan(&op)

	// fmt.Println("Enter second operator")
	// fmt.Scan(&num2)

	// result := 0

	// switch op {
	// case "+":
	// 	result = num1 + num2
	// case "-":
	// 	result = num1 - num2
	// case "*":
	// 	result = num1 * num2
	// case "/":
	// 	if num2 == 0 {
	// 		fmt.Println("Error! division by zero")
	// 		return
	// 	}
	// 	result = num1 / num2
	// }

	// fmt.Println("Result:",result)

