package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app!")
	fmt.Println("Please rate our pizza between 1 to 10:")

	reader := bufio.NewReader(os.Stdin)
	//	reader.ReadString('\n')

	input, _ := reader.ReadString('\n')
	fmt.Println("You rated to out pizza: ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your  rating: ", numRating+1)
	}
}
