package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cant devide by zero")
	}
	return a / b, nil
}

func main() {
	result, err := divide(78, 8)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("result:", result)
}
