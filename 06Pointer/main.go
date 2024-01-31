package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of Pointers")
	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 25
	var ptr = &myNumber
	fmt.Println("Value of pointer is ", ptr)
	fmt.Println("Value of myNumber is ", *ptr)
	*ptr =  *ptr * 3
	fmt.Println("Value of new number is ", *ptr)
	fmt.Println("Value of new myNumber is ", myNumber)

}
