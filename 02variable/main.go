package main

import "fmt"


const LoginToken = "HLLEOeekjodfKJKNFJmsdjf32349320"  //public

func main() {
	var username string = "Jack"
	fmt.Println("Your name is :" + username)
	fmt.Printf("variable is of type: %T \n", username)

	var IsLoggedIn bool = false
	fmt.Println(IsLoggedIn)
	fmt.Printf("variable is of type: %T \n", IsLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	//default value and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("variable is of type: %T \n", anotherString)

	var anotherfloat float64
	fmt.Println(anotherfloat)
	fmt.Printf("variable is of type: %T \n", anotherfloat)

	var anotherBool bool
	fmt.Println(anotherBool)
	fmt.Printf("variable is of type: %T \n", anotherBool)

	//implicit type
	var website ="www.google.com"
	fmt.Println(website)
	fmt.Printf("variable is of type: %T \n", website)

	//no var style
	numberOfUsert := 235.65
	fmt.Println(numberOfUsert)
	fmt.Printf("variable is of type: %T \n", numberOfUsert)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)

}
