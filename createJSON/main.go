package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //this will be an alias for name. when using this struct for json, it'll replace name with coursename
	Price    int
	Platform string   `json:"website "`
	Password string   `json:"-"`              //providing a - will make this field disapper from api request. this provides security
	Tags     []string `json:"tags,omitempty"` //omitempty will remove the fields when value is null #note:no space between tags,omitempty
}

func main() {
	fmt.Println("welcome to creating JSON data")
	// EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	lcoCourses := []course{{"reactjs", 500, "somereactwebsite", "somepassword1", []string{"web-dev", "js"}},
		{"mern", 600, "somemernwebsite", "somepassword2", []string{"web-dev", "js", "express", "mongo"}},
		{"python", 700, "somepywebsite", "somepassword3", nil},
	}

	finaljson, _ := json.MarshalIndent(lcoCourses, "", "\t")
	fmt.Printf("%s\n", finaljson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "mern",
		"Price": 600,
		"website ": "somemernwebsite",
		"tags": ["web-dev", "js", "express", "mongo"]
	}
	`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("VALID JSON ")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse) //The %#v variant prints a Go syntax representation of the value, i.e. the source code snippet that would produce that value.
	} else {
		fmt.Println("INVALID JSON")
	}

	//some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Println("JSON DATA USING MAP INTERFACE:")
	fmt.Printf("%#v\n", myOnlineData)
	fmt.Println("PRINTING IT NEATLY USING FOR LOOP:")

	for k, v := range myOnlineData {
		fmt.Printf("key: %s, value: %v, type: %T\n", k, v, v)
	}

}
