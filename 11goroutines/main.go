package main

// import "fmt"

import (
	"fmt"
	"time"
)

// func greet() {
// 	fmt.Println("from goroutines...")
// }

func main() {
	go func() {
		fmt.Println("from goroutines...")
	}()
	fmt.Println("from main....")

	time.Sleep(1 * time.Second)
}

// interface

// type Speaker interface {
// 	speak()
// }

// type Dog struct{}
// type Person struct{}

// func (d Dog) speak() {
// 	fmt.Println("Woof!")
// }

// func (p Person) speak() {
// 	fmt.Println("Hey!")
// }

// func main() {
// 	var s Speaker
// 	s = Dog{}
// 	s.speak()

// 	s = Person{}
// 	s.speak()
// }

// func describe(i interface{}) {
// 	fmt.Printf("Value: %v, Type: %T\n", i, i)
// }

// func main() {
// 	describe(43)
// 	describe("fazal")

// }
