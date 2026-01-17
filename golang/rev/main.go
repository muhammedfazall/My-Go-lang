package main

import (
	"fmt"
	"strings"
)

// func cycleNames(n []string, f func(string)) {
// 	for _, v := range n {
// 		f(v)
// 	}
// }

// func greet(n string) {
// 	fmt.Printf("hey %v\n", n)
// }


func getInitials(n string) (string,string) {
	s := strings.ToUpper(n)
	names := strings.Split(s," ")

	var initials []string

	for _,v := range names{
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0],initials[1]
	}
	return initials[0] , ""
}

func main() {

	a1,a2 := getInitials("fazal muhammed")
	b1,b2 := getInitials("fazal")

	// names := []string{"faz", "laz", "zal", "aza"}
	// sort.Strings(names)
	fmt.Println(a1,a2)
	fmt.Println(b1,b2)

	// cycleNames(names, greet)
}
