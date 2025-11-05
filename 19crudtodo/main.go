package main

import (
	"fmt"
	"time"
)

func main() {

	choice := 0
	todo := map[int]string{
		1: "work hard",
		2: "level up",
		3: "escape",
	}
	nextId := 4

	for {
		fmt.Println("To-do-list")
		fmt.Println("-----------------------------------")
		fmt.Println("1. view your To-do-list")
		fmt.Println("2. Add Task")
		fmt.Println("3. Delete a Task")
		fmt.Println("4. update a task")
		fmt.Println("0. Exit")
		fmt.Println("Enter your choice: ")
		fmt.Scan(&choice)
		if choice == 0 {
			break
		}
		switch choice {
		case 1:
			fmt.Println("To-Do-List")
			fmt.Println("-----------------------------")
			i := 1
			for key, value := range todo {
				fmt.Printf("%d- id=%d: %s\n", i, key, value)
				i++
			}
		case 2:
			fmt.Scanln()
			var new string
			fmt.Scanln(&new)
			todo[nextId] = new
			nextId++
		case 3:
			fmt.Println("chose the task to delete:")
			deleteId := 0
			fmt.Scan(&deleteId)
			delete(todo, deleteId)
		case 4:
			fmt.Println("chose the task to update:")
			updateId, new := 0, ""
			fmt.Scan(&updateId)
			fmt.Scan(&new)
			todo[updateId] = new
		default:
			fmt.Println("chose u fool")
		}
		time.Sleep(time.Second * 2)
	}
	fmt.Println("Thankyou!")
}
