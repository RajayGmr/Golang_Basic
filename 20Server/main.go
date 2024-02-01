package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to a class of Creating server")
	PerformGetRequest()
	PerformPostRequest()
	PerfromPostFormRequest()

}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is: ", byteCount)
	fmt.Println("Get Method:", responseString.String())

	//fmt.Println(content)
	//fmt.Println(string(content))
}

func PerformPostRequest() {
	const myurl1 = "http://localhost:8000/post"

	//fake json payload
	requestBody := strings.NewReader(`
	{
        "coursename":"Let's go with golang",
		"price":0,
		"platform":"learnfromsomewhere.edu"
    }`)
	response, err := http.Post(myurl1, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println("Post Method: ", string(content))
}

func PerfromPostFormRequest() {
	const myurl2 = "http://localhost:8000/postform"

	//formData
	data := url.Values{}
	data.Add("firstname", "Jack")
	data.Add("Lastname", "Smith")
	data.Add("email", "jacksmith@gmail.com")

	response, err := http.PostForm(myurl2, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println("Post Form Method: ", string(content))
}
