package main

import "fmt"

func main() {
	var a int

	for {

		fmt.Println("Enter the number: ")
		fmt.Scan(&a)

		if a%2 == 0 {
			fmt.Printf("The number %v is even\n", a)
		} else {
			fmt.Printf("The number %v is odd\n", a)
		}
	}
}
