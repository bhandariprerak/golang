package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:9077/learn?coursename=reactjs&paymentid=12345"

func main() {
	fmt.Println("welcome to handling urls in go")

	//parsing
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println("result scheme:", result.Scheme)
	fmt.Println("result host:", result.Host)
	fmt.Println("result path :", result.Path)
	fmt.Println("result raw query:", result.RawQuery)
	fmt.Println("result port :", result.Port())

	qparams := result.Query() //the type of qparams is url.Values. So, qparams will have key-value structure.
	fmt.Println("qparams", qparams)
	fmt.Println("the individual queries are:")
	for _, val := range qparams {
		fmt.Println("the values of qparams are:", val)
	}

	// to create a url always pass a refernece, not a copy. so use &url

	partsOfUrl := &url.URL{ //note: use &url here
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/prerak",
		RawQuery: "a=12&b=45",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("url created from parts:", anotherUrl)

}
