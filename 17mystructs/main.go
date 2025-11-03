package main

import "fmt"

func main() {

	fazal := User{"Fazal", 20, "fake@fake.com", true}
	fmt.Printf("Users detail: %+v\n", fazal)
	fmt.Printf("%v is %v years of old\n", fazal.Name, fazal.Age)

	fazal.getStatus()
	
}

type User struct {
	Name   string
	Age    int
	Email  string
	Status bool
}

func (u User) getStatus(){
	fmt.Println("The User is online? :",u.Status)
}