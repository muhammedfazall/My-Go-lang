package main

import "fmt"

func divide(n1 int, n2 int) int {
	if n2 == 0 {
		panic("Cant divide by zero")
	}
	return n1 / n2
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic Recovered:", r)
		}
	}()

	fmt.Println(divide(48, 5))
	fmt.Println(divide(50, 0))

}
