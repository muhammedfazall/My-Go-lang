package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("panic recovered: %v", r)
		}
	}()

	fmt.Println("before panic")
	panic("error occured")
	fmt.Printf("after panic")
}
