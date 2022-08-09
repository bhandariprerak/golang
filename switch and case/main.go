package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("welcome to switch and case. here we demonstrate a simple ludo game dice roll.")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Printf("value of dice is: %v\n", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spaces")
	case 3:
		fmt.Println("You can move 3 spaces")
		fallthrough
	case 4:
		fmt.Println("You can move 4 spaces")
	case 5:
		fmt.Println("You can move 5 spaces")
	case 6:
		fmt.Println("You can move 6 spaces and roll again")
	default:
		fmt.Println("What was that?!")
	}
}

// fallthrough will execute the next case also instead of breaking
