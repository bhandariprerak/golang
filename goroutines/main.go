package main

import (
	"fmt"
	"net/http"
	"sync"
)

//storing the signals
var signals = []string{"test"}

//waitgroup declaration for goroutines
var wg sync.WaitGroup //usually these are pointers

var mutex sync.Mutex //usually these are pointers

func main() {
	// go greeter("Hello")
	// greeter("World")

	websitelist := []string{
		"https://google.com",
		"https://fb.com",
		"https://github.com",
		"https://go.dev",
		"https://instagram.com",
		"https://gooogle.com",
		"https://lco.dev",
	}

	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait() //this usually always goes at end of main method
	fmt.Println("signals:", signals)
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(2 * time.Millisecond)
// 		fmt.Println(i+1, s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done() //this indicates work done
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error in fetching status code")
	} else {
		mutex.Lock()
		signals = append(signals, endpoint)
		mutex.Unlock()
		fmt.Printf("%d: status code for %s\n", res.StatusCode, endpoint)

	}
}
