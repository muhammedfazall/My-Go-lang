package main

import "fmt"

func fibonacci(n int) {
	a, b := 0, 1
	// for i := 0; i < n; i++ {  // for first n fibonacci numbers
	for a < n{  // for fibonacci numbers upto n
		fmt.Printf("%d ",a)
		a,b = b,a+b
	}
	fmt.Println()
}