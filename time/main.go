package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of go lang")
	currTime := time.Now()
	fmt.Println("Current time is:")
	fmt.Println(currTime)
	fmt.Println("Formatted time is:")
	fmt.Println(currTime.Format("01-02-2006 Monday 15:04:06")) //format is standard for every conversion.
	fmt.Println("create time format:")
	createdDate := time.Date(2022, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}
