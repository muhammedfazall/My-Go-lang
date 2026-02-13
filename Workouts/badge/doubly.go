package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) AddtoFront(data int) {
	newNode := &Node{
		data: data,
		next: l.head,
		prev: nil,
	}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	l.head.prev = newNode
	l.head = newNode
}

func (l *LinkedList) AddtoEnd(data int) {
	newNode := &Node{
		data: data,
		next: nil,
		prev: l.tail,
	}

	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	l.tail.next = newNode
	l.tail = newNode
}

func (l *LinkedList) Display() {
	for cur := l.head; cur != nil; cur = cur.next {
		fmt.Printf(" %d ", cur.data)
	}
}

// func main() {
// 	l := LinkedList{}

// 	fmt.Println("forward")
// 	l.AddtoFront(10)
// 	l.AddtoFront(20)
// 	l.AddtoFront(30)
// 	fmt.Println("back")
// 	l.AddtoEnd(30)
// 	l.AddtoEnd(20)
// 	l.AddtoEnd(10)

// 	l.Display()

// }

// // type Stack struct{
// // 	data []int
// // }

// // func (s *Stack) Push(v int)  {
// // 	s.data = append(s.data, v)
// // }
// // func (s *Stack) Pop() (int,error) {
// // 	if len(s.data) == 0 {
// // 		return 0,errors.New("empty")
// // 	}
// // 	i:= len(s.data)-1
// // 	v := s.data[len(s.data)-1]
// // 	s.data = s.data[:i]
// // 	return v, nil
// // }

// // func (s *Stack) Display() ([]int,error) {
// // 	if len(s.data) == 0 {
// // 		return nil,errors.New("empty")
// // 	}
// // 	return s.data,nil
// // }

// // func (s *Stack) Push(v int)  {

// // }

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// type Slice struct {
// 	data []int
// }

// func (s *Slice) Enqueue(v int) {
// 	s.data = append(s.data, v)
// }

// func (s *Slice) Dequeue() (int, error) {
// 	if len(s.data) == 0 {
// 		return 0,errors.New("empty")
// 	}

// 	v := s.data[0]
// 	s.data = s.data[1:]
// 	return v,nil
// }

// func main() {
// 	s := Slice{}

// 	s.Enqueue(10)
// 	s.Enqueue(20)
// 	s.Enqueue(30)

// 	fmt.Println(s.data)
// 	v,_ := s.Dequeue()
// 	fmt.Println(v)
// }
