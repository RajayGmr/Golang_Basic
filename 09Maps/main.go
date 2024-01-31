package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of Maps")
	language := make(map[string]string)

	language["Go"] = "Golang"
	language["Java"] = "Java"
	language["Python"] = "Python"
	language["C++"] = "C plus plus"
	language["Js"] = "JavaScript"

	fmt.Println("List of all languages=>", language)
	fmt.Println("Js Shorts for:", language["Js"])

	delete(language, "Java")
	fmt.Println("List of all languages=>", language)

	//loop in golang 
	for key,value:=range language{
		fmt.Printf("For key: %v, value is %v\n", key, value)
	}

}
