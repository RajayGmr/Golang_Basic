package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of function")
	greeter()
	greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is:", result)

	resultTwo := stringjoin("You are", "amazing",)
	fmt.Println("Result is:", resultTwo)

	proResult, Mymessage := proAdder(1,2,3,4,5,6,7,8,9,10)
	fmt.Println("Result is:", proResult)
	fmt.Println("Message is:", Mymessage)

}

func greeter() {
	fmt.Println("Namstey from golang")
}

func greeterTwo() {
	fmt.Println("Another metod")
}
func adder(n int, m int) int {
	return n + m
}
func stringjoin(s string, m string) string{
	return s +" "+ m
}

func proAdder(values ...int) (int,string) {
	sum := 0
    for _, value := range values{
        sum += value
    }
    return sum, "Hi Pro result function"
}
