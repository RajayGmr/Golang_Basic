package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of Methods")

	jack := User{"jack", "jack342@gmail.com", true, 25}
	fmt.Println(jack)

	fmt.Printf("Jack details are %+v\n", jack)
	fmt.Printf("Jack details are %v and email is %v\n", jack.Name, jack.Email)
	fmt.Printf("Jack details are %T and email is %T\n", jack.Name, jack.Email)
	fmt.Printf("Jack details are %s and email is %s\n", jack.Name, jack.Email)

	jack.GetStatus()
	jack.NewMail()
	fmt.Printf("Jack details are %v and email is %v\n", jack.Name, jack.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User active:", u.Status)
}
func (u User) NewMail() {
	u.Email = "James@gmail.com"
	fmt.Println("Email for this user is:", u.Email)
}
