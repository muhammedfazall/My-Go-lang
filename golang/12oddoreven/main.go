// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("Check if Odd or Even")

// 	for { //non-stop
// 		num := 0
// 		yn := ""
// 		fmt.Println("Enter a Number:")
// 		fmt.Scan(&num)

// 		if num%2 == 0 {
// 			fmt.Println("Even")
// 		} else{
// 			fmt.Println("Odd")
// 		}	
// 		fmt.Println("You wanna continue ? Y / N")
// 		fmt.Scan(&yn)
// 		if yn == "n" || yn == "N" {
// 			break
// 		}
// 	}
// }


package main

import "fmt"

func main() {
	n := 10 // how many Fibonacci numbers you want
	a, b := 0, 1

	fmt.Print("Fibonacci Sequence: ")

	for range n {
		fmt.Print(a, " ")
		next := a + b
		a = b
		b = next
	}
}

