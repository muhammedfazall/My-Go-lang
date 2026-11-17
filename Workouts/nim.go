package main

import "fmt"

func main() {

	// fmt.Println(sum(20,30))
	// fmt.Println(sum(50,30,1,42,55))
	// fmt.Println(sum(20,30,88,455))

	c := counter()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())
}

func sum(numbers ...int) int {
	res := 0

	for _, num := range numbers {
		res += num
	}

	return res
}

func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}
}