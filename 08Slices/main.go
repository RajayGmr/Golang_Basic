package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to a class of Slices...")

	var fruitList = []string{"apple", "orange", "peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "coconut", "Mango")
	fmt.Println("Fruit list are:", fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println("Fruitlist are:", fruitList)

	highScore := make([]int, 5)
	highScore[0] = 250
	highScore[1] = 100
	highScore[2] = 50
	highScore[3] = 200
	highScore[4] = 300
	highScore = append(highScore, 444, 900, 700)
	fmt.Println(highScore)

	fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))

	//how to remove a value from slices based on index 
	var course =[]string{"Vuejs", "java","python","goalng","C++"}
	fmt.Println("List of course are==>", course)
	var index int =2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)
}
