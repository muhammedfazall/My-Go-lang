package main

import "fmt"

func sort(arr []int) {
	for i := range arr {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func main() {
	var arr = []int{4, 3, 2, 1}

	fmt.Println(arr)
	sort(arr)

	fmt.Println(arr)

	fmt.Println(len(arr))
}
