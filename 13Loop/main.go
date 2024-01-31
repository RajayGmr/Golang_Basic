package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of loop")
	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}
	fmt.Println("--------------------------------")

	for i := range days {
		fmt.Println(days[i])
	}
	fmt.Println("--------------------------------")

	for index, day := range days {
		fmt.Println(index, day)
	}
	fmt.Println("--------------------------------")

	rougueValue := 1
	for rougueValue < 10 {
		if rougueValue == 3 {
			goto lco
		}
		if rougueValue == 6 {
			break
		}
		fmt.Println("Value is:", rougueValue)
		rougueValue++

	}
	lco:
	fmt.Println("Jump to www.google.com")
}
