package main

import "fmt"

func main() {
	// student := map[string]int{
	// }

	// student := make(map[string]int)

	var student = make(map[string]int)

	student["afzah"] = 25
	student["awjhd"] = 25
	fmt.Println(student)
	
	student["aliue"] = 54
	student["afzah"] = 65
	
	fmt.Println(student)

	for key, value := range student{
		fmt.Printf("key: %v value: %v\n",key, value)
	}
}
