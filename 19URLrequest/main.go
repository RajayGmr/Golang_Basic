package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to the class of URL Request")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query parameters are: %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams{
		fmt.Println("Param is:",val)
	}
	partofUrl := &url.URL{
		Scheme:   result.Scheme,
        Host:     result.Host,
        Path:     result.Path,
        RawQuery: result.RawQuery,
	}
	anotherUrl := partofUrl.String()
	fmt.Println(anotherUrl)
}
