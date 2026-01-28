package main

func Divide(a,b int) int {
	if b == 0 {
		panic("Cant divide by zero")
	}
	return a/b
}