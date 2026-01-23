package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func ()  {
		fmt.Println("one")
		wg.Done()	
	}()

	go func ()  {
		fmt.Println("two")
		wg.Done()	
	}()

	wg.Wait()
	fmt.Println("finsised")
}