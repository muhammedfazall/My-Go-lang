package main

import "fmt"

type Person struct {
	name string
	age  int
	city string
}

func (p Person) isAdult() bool {
	return p.age >= 18
}

func (p *Person) decAge() {
	p.age = p.age - 10
}

// func (p Person) greet() {
// 	fmt.Printf("hey %s, you are %d years old now!", p.name, p.age)
// }

func main() {
	p := Person{name: "fazal", age: 24}
	// p.greet()
	fmt.Println(p.isAdult())
	p.decAge()

	fmt.Println(p.isAdult())
	fmt.Println(p.age)
}

// var p1 Person
// p1.name = "fazal"
// p1.age = 24
// p1.city = "calicut"

// fmt.Println(p1)
// fmt.Println(p1.age)
