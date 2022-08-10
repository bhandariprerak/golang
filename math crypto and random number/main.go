package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
	// "time"
)

func main() {
	fmt.Println("welcome to math, crypto and random numbers in go")
	// var mynum1 int = 2
	// var mynum2 float64 = 4.56

	// fmt.Println("the sum of both numbers is:", mynum1+int(mynum2)) //this gives error

	//random number from math/rand
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println("random number generated from math/rand is:", rand.Intn(6))

	// random number from crypto/rand
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println("random number generated from crypto/rand is:", myRandomNum)

}
