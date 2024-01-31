package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of Defer")
	defer fmt.Println("Please pay attention")
	myDefer()
}

func myDefer() {
	for i := 0; i < 6; i++ {
		defer fmt.Println(i)
	}
}
