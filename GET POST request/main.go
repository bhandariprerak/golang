package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("welcome to get requests in go")
	PerformGetRequest()
	PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("status code of request is:", response.Status)
	fmt.Println("content length is:", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	// 2 ways to read and print the data content:
	//method 1: using strings library and using builder, write and string()
	//this is lengthy but more powerful as it provides more methods and it is a library
	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)
	fmt.Println("bytecount is:", byteCount)
	fmt.Println("response string is:", responseString.String())

	//method 2: using string wrapper to directly convert raw data to string.
	//this is easy method but less powerful
	// fmt.Println("content is:", string(content))

}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
	  {
		 "coursename" : "learn-golang",
		 "price" : "priceless",
		 "platform" : "by yourself"
	  }
	`)
	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("the content POSTed is:")
	fmt.Println(string(content))

}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"
	data := url.Values{}
	data.Add("firstname", "prerak")
	data.Add("lastname", "bhandari")
	data.Add("email ", "prerak@gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("the data from PostForm:")
	fmt.Println(string(content))

}
