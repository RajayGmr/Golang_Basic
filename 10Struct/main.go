package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of Struct")

	Jack := User{"Jack", "jack24@gmail.com", true, 23}
	fmt.Println(Jack)
	fmt.Printf("Jack details are %+v\n", Jack)
	fmt.Printf("Name is %v and email is %v.",Jack.Name,Jack.Email)

	//no inheritance in golang  // no super or parent 
}
type User struct {
	Name string
	Email string
	Status bool
	Age int
}