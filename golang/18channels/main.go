package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	fmt.Println("hety")

	go SendDataToChannel(ch, 145)

	v := <-ch

	fmt.Println("value passed:", v)

	fmt.Println("----------------------------------------")

	// sending custom data through channel

	p:=Person{"fazal","kkd",7}

	pch := make(chan Person)

	go SendPerson(pch,p)

	place := (<- pch).Place
	fmt.Println(place)

}

func SendDataToChannel(ch chan int, value int) {
	ch <- value
}

type Person struct{
	Name string
	Place string
	Id int
}

func SendPerson(ch chan Person,p Person ){
	ch<-p
}