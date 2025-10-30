// Assigning an Anonymous Function to a Variable

// package main

// import "fmt"

// func main() {
// 	greet := func(name string) string {
// 		return "hey " + name
// 	}
// 	message := greet("fazal")

// 	fmt.Println(message)
// }

// Anonymous Function with Parameters

// package main

// import "fmt"

// func main() {
// 	func(name string){
// 		fmt.Printf("hey %v", name)
// 	}("fazal")
// }

//passing an anonymous fn as an argument

// package main

// import "fmt"

// func process(fn func(int, int) int) {
// 	result := fn(10, 25)
// 	fmt.Printf("Result: %v", result)
// }

// func main() {
// 	process(func(a, b int) int {
// 		return a + b
// 	})
// }

// package main

// import "fmt"

// func process(fn func(int, int, int) float32) {
// 	result := fn(10, 25, 2)
// 	fmt.Printf("Result: %v", result)
// }

// func main() {
// 	process(func(a, b, c int) float32 {
// 		return float32(a + b) / float32(c)
// 	})
// }

//closures : The anonymous function "remembers" the variable counter even though
//           it’s outside the function — this is called a closure.

// package main

// import "fmt"

// func main() {

// 	counter := 0

// 	increment := func() int {
// 		counter++
// 		return counter
// 	}

// 	fmt.Println(increment())
// 	fmt.Println(increment())
// 	fmt.Println(increment())
// }

package main

import "fmt"

func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func main() {
	double := multiplier(2)
	triple := multiplier(3)
	quadruple := multiplier(4)

	fmt.Println(double(16))
	fmt.Println(triple(3))
	fmt.Println(quadruple(5))
}
