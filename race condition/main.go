package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("welcome to race condition in go")

	//wait group for go routines
	wg := &sync.WaitGroup{}

	//mutex
	// mut := &sync.Mutex{}
	mut := &sync.RWMutex{}

	var score = []int{0}

	wg.Add(3)
	// lambda functions - goroutines
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Four R")
		mut.RLock()
		fmt.Println("score:", score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println("score:", score)
}
