package main

import (
	"fmt"
	"time"
)

func greeet(n string) {
	fmt.Println("hey", n)
}

func main() {
	go func() {
		fmt.Println("hey go")
	}()

	go func() {
		fmt.Println("hey n")
	}()

	go greeet("fazal")

	fmt.Println("hey main")

	time.Sleep(2 * time.Second)
}
