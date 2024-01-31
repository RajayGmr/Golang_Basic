package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to a class of File")
	content := "this need to go in a file - www.google.com\n"

	file, err := os.Create("./myfile.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d bytes\n", length)
	defer file.Close()

	readFile("./myfile.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read %s\n", filename)
	fmt.Println("Text data inside the file =>", string(databyte))

}
