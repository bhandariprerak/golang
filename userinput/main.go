package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user inputs"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter rating for our service: ")

	// comma ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the input")
	fmt.Printf("Type of this rating is : %T \n", input)

}
