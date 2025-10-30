package main

import "fmt"

func main() {

	var age int = 22
	p := &age
	fmt.Println(age)
	fmt.Println(p)
	fmt.Println(*p)
}
