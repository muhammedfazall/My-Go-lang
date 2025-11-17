package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic recovered", r)
		}
	}()

	fmt.Println(divide(50,0))

}

func divide(n1, n2 int) int {
	if n2 == 0 {
		panic("cant divide with zero")
	}
	return n1 / n2
}
