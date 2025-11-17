package main

import (
	"errors"
	"fmt"
)

func main() {
	var name string = "fazal"
	printMe(name)
	result, reminder, err := intDivision(59, 5)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Result: %v\nReminder: %v", result, reminder)
}
func printMe(name string) {
	fmt.Printf("hey hey %v \n", name)
}
func intDivision(numerator int, denomenator int) (int, int, error) {
	if denomenator == 0 {
		err := errors.New("cant divide by 0")
		return 0, 0, err
	}

	var result = numerator / denomenator
	var reminder = numerator % denomenator
	return result, reminder, nil
}
