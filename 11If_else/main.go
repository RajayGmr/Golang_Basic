package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of IF Else")
	logInCount := 25
	var result string

	if logInCount < 10 {
		result = "Regulare User"
	}else if logInCount>10 {
		result = "Average User"
    }else{
		result = "Premium User"
	}
	fmt.Println(result)

	if 9%2==0{
		fmt.Println("number is even")
	}else{
		fmt.Println("number is odd")
	}

	if num:= 5; num<=9{
		fmt.Println("number is less than 10")
	}else{
		fmt.Println("number is greater than 10")
	}

}
