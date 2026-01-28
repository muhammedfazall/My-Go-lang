package main

import "fmt"

func oddOrEven(a int) {
	if a%2 == 0 {
		fmt.Printf("The number %v is even\n", a)
	} else {
		fmt.Printf("The number %v is odd\n", a)
	}
}
