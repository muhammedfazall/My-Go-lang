package main

import "fmt"

func main() {
	newArr := [3]int32{1, 2, 3}
	fmt.Println(newArr[0])
	newArr[2] = 123
	fmt.Println(newArr)
	fmt.Println(newArr[1:3])

	newSlice := []int32{4, 5, 7}
	fmt.Println(len(newSlice),cap(newSlice))

	newSlice = append(newSlice, 7,8)
	fmt.Println(newSlice,len(newSlice),cap(newSlice))


	newSlice2 := make([]int32, 3, 8)
	fmt.Println(newSlice2, cap(newSlice2))
}
