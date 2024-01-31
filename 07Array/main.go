package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of Array")

	var fruitList [4]string
	// fruitList = append(fruitList, "Apple")
	// fruitList = append(fruitList, "Orange")
	// fmt.Println(fruitList)

	fruitList[0]="Apple"
	fruitList[1]="Orange"
	fruitList[3]="grape"
	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list array len is: ",len(fruitList))

	var vegList = [3]string{"beans","potatoes","tomatoes"}
	fmt.Println("Veg list is: ", vegList)
	fmt.Println("Veg list array len is: ",len(vegList))
}
