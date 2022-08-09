package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("welcome to web requests in go")
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("response is of type: %T\n", response)
	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databytes))
}
