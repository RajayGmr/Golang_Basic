package main

import (
	"fmt"
	"io"
	"net/http"
)

const url= "https://loc.dev"

func main() {
	fmt.Println("Welcome to a class of WEB Requests.")

	response, err :=http.Get(url)
	if err!= nil {
        fmt.Println(err)
    }
	fmt.Printf("Type of response: %T\n", response)
	fmt.Println(response)
	defer response.Body.Close() //caller's responsibilityto close the connection

	databyte, err := io.ReadAll(response.Body)
	if err!= nil {
        fmt.Println(err)
    }
	content := string(databyte)
	fmt.Println(content)
}
