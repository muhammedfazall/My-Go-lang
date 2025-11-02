package main

import "fmt"

func main() {
	fruitList := []string{"apple", "mango", "orange", "peach", "pineapple"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fmt.Println(fruitList)
	fmt.Println(fruitList[1:])

	// to remove a value from slice using index:

	var veg = []string{"carrot", "potato", "onion", "beetroot", "pumpkin", "brinjal"}
	fmt.Println(veg)
	var index int = 2
	fmt.Println(append(veg[:index], veg[index+1:]...)) //ellipsis: you can "spread" its elements as 'individual arguments' by appending ... to the slice name


}
