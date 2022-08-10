package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("welcome to channels in golang")

	myCh := make(chan int, 2) //the 2nd params defines how many times this channel will be used.

	// myCh <- 5 //to insert a value in myCh. in golang we only use  "<-" arrow. never "->" arrow.
	// fmt.Println(<-myCh)

	//note: channels in go require someone to listen to that channel and serve to the channel at the same time.

	wg := &sync.WaitGroup{}
	wg.Add(2)
	// we create 2 channels atleast because:
	// First one will be responsible for receiving the values
	// Second one will be responsible for putting values in the chaneel
	go func(ch <-chan int, wg *sync.WaitGroup) { // adding <- before chan makes this channel RECIEVE ONLY
		// close(myCh) //now this gives error because of <- in params.
		val, isChannelOpen := <-myCh
		fmt.Println("val:", val, "channel status:", isChannelOpen)
		// fmt.Println(<-myCh)
		// fmt.Println(<-myCh) // we can put listeners in for loop
		wg.Done()
	}(myCh, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) { //adding <- after chan makes this channel SEND ONLY
		// close(myCh)
		myCh <- 0
		myCh <- 6   //if we add this channel as well then we need 2 listeners.
		close(myCh) //closing the channel after values are put. close only after using the channels
		wg.Done()
	}(myCh, wg)
	wg.Wait()
}

//if channel is closed before using it, it gives 0 as output.
//so we can get confused if we actually did myCh <- 0 or it's error.
//to know that we use val, isChannelOpen := <- myCh
