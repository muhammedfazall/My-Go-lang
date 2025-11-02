package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)

	fmt.Println(languages["RB"])

	delete(languages, "PY")

	fmt.Println(languages)

	// iterate over map

	for key, value := range languages{
		fmt.Printf("for key %v , value is %v\n", key, value)
	}

	// if key not needed: for _,value := ....... 
}
