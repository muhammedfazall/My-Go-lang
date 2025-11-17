package main

// import "fmt"

import (
	"fmt"
	"time"
)

// func greet() {
// 	fmt.Println("from goroutines...")
// }

func main() {
	go func() {
		fmt.Println("from goroutines...")
	}()
	fmt.Println("from main....")

	time.Sleep(1 * time.Second)
}


