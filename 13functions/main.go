package main

import "fmt"

func calculation(num1, num2 float64) (sum, product, quo, diff float64) {
	sum = num1 + num2
	product = num1 * num2
	quo = num1 / num2
	diff = num1 - num2

	return sum, product, quo, diff
}
func main() {
	a := 5.0
	b := 10.0

	s, p, q, d := calculation(a, b)

	fmt.Println(":", s, p, q, d)
}
