package main

import "fmt"

func main() {
	// var num int
	// fmt.Println("Enter the number: ")
	// fmt.Scan(&num)

	// oddOrEven(num)
	
	fibonacci(120)

	defer func ()  {
		if r := recover() ; r != nil{
			fmt.Println("panic recovered",r)
		}
	}()

	result := Divide(2,0)
	fmt.Println(result)
}
