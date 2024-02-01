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
	DecodeJson()
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

func DecodeJson (){
	jsonDataFromweb := []byte(`
		{
			"coursenames":"ReactJs",
			"price":1200,
			"Website":"https://reactjs.org/",
			"tags":[
                "ReactJs",
                "Javascript",
                "React"
            ]
		}
	`)
	var MyCourse course
	checkValid := json.Valid(jsonDataFromweb)
	if checkValid {
		fmt.Println("JSon was valid")
		json.Unmarshal(jsonDataFromweb, &MyCourse)
		fmt.Print("%#v\n", MyCourse)
	}else{
		fmt.Println("JSon was not valid")
	}

	//some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromweb, &myOnlineData)
	fmt.Printf("my Online Data :%#v\n", myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", key, value, value)
	}
}
