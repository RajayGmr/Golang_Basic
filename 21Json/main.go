package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int 
	platform string `json:"Website"`
	Password string `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to a class of Json")
	EncodeJson()
}

func EncodeJson() {
	MyCourse := []course{
		{
			"Java",
			1200,
			"Android",
			"2344",
			[]string{"Java", "Android", "nil"},
		},
		{
			"C++",
			800,
			"Android",
			"890324",
			[]string{"C++", "Android", "Web"},
		},
		{
			"Python",
			1000,
			"Android",
			"sdKxY32334",
			[]string{"Python", "Android", "Web"},
		},
	}
	//package this data as Json data
	jsonData, err := json.MarshalIndent(MyCourse, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", jsonData)
	fmt.Println("--------------------------------")
	fmt.Println(string(jsonData))
}
