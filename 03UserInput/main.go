package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to User Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the rating 1-10 for our Pizza :")

	//comma ok // err ok syntax
	input, err := reader.ReadString('\n')
	fmt.Println("You rated our pizza: " + input)
	fmt.Printf("variable is of type: %T \n", input)
	if err != nil {
		fmt.Println(err)
	}
}
