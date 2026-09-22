package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {

	// s := "heyfazal"

	// r := []byte(s)
	// fmt.Println("hey world")

	// result := sum.Sum(4,5)

	// for i := 0 ; i < len(s); i++ {
	// 	fmt.Printf("%v ",s[i])
	// }

	// for _,v := range r{

	// 	v := string(v)
	// 	fmt.Printf("%v ",v)
	// }

	ar := findWordsContaining([]string{"abc", "bcd", "aaaa", "cbc"}, 'a')

	fmt.Println(ar)

}

func findWordsContaining(words []string, x byte) []int {
	indices := []int{}
	for i, w := range words {
		for _, v := range w {
			if byte(v) == x {
				indices = append(indices, i)
			}
		}
	}
	return indices
}

func isValid(s string) bool {
	if len(s) > 5 {
		return false
	}

	if len(strings.TrimSpace(s)) != len(s) {
		return false
	}

	for _, ch := range s {
		if !unicode.IsDigit(ch) {
			return false
		}
	}

	return true
}
