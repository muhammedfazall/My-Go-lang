package main

// import "fmt"

func main() {

	// var num int
	// fmt.Println("Enter the number: ")
	// fmt.Scan(&num)

	// oddOrEven(num)

	// fibonacci(120)

	// defer func ()  {
	// 	if r := recover() ; r != nil{
	// 		fmt.Println("panic recovered",r)
	// 	}
	// }()

	// result := Divide(2,0)
	// fmt.Println(result)

	//   fmt.Println(Count(-11651))
	//   fmt.Println(Count(15411651))

	//   fmt.Println(IsClear(4561))

}

func mostWordsFound(sentences []string) int {
	wordCount := map[string]int{}
	for _,s := range sentences {
		count := 0
		for _,ch := range s {
			if string(ch) == " " {
				count++
			}
		}
		wordCount[s] = count + 1
	}

	max := 0
	for v := range wordCount {
		if wordCount[v] > max {
			max = wordCount[v]
		}
	}
	return max
}
