package main

import "fmt"

// interface

type Speaker interface {
	speak()
}

type Dog struct{}
type Person struct{}

func (d Dog) speak() {
	fmt.Println("Woof!")
}

func (p Person) speak() {
	fmt.Println("Hey!")
}

func describe(i interface{}) {
	fmt.Printf("Value: %v, Type: %T\n", i, i)
}

func main() {
	var s Speaker
	s = Dog{}
	s.speak()

	s = Person{}
	s.speak()
	
	describe(43)
	describe("fazal")
}
