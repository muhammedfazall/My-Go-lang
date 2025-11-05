package main

import "fmt"

func main() {

	//goroutine

	go func (){
		fmt.Println("hello from goroutine!")
	}()

	//struct

	type Person struct{
		Name string
		Age int
	}

	fazal := Person{"Fazal",24}
	fmt.Printf("Name: %v| Age: %v\n",fazal.Name,fazal.Age)

	// map

	students := map[string]int{
		"fazal":24,
		"anandu":24,
	}
	fmt.Println(students["fazal"])

	//slice

	mySlice := []int{2,3,5,8}
	fmt.Println(mySlice)
}
