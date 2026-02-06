package main

import (
	"strconv"
)

func Count(n int) int {
	for n < 0 {
		n = -n
	}
	return len(strconv.Itoa(n))
}

// func Count(n int) int {
//   if n == 0 {
//     return 1
//   }

//   if n < 0 {
//     n = -n
//   }

//   count := 0
//   for n > 0 {
//   n /= 10
//   count ++
//   }

//   return count
// }

func IsClear(n int) bool {
	sum := 0
	temp := n

	for temp > 0 {
		sum += temp % 10
		temp /= 10
	}

	return n%sum == 0
}
