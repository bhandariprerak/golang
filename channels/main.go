package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("welcome to channels in golang")

	myCh := make(chan int)

	// myCh <- 5 //to insert a value in myCh. in golang we only use  "<-" arrow. never "->" arrow.
	// fmt.Println(<-myCh)

	//note: channels in go require someone to listen to that channel and serve to the channel at the same time.

	wg := &sync.WaitGroup{}
	wg.Add(2)
	// we create 2 channels atleast because:
	// First one will be responsible for receiving the values
	// Second one will be responsible for putting values in the chaneel
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		myCh <- 5
		wg.Done()
	}(myCh, wg)

}
