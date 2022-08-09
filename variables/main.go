package main

import "fmt"

//num:=34 will give an error here

const FirstLetterCapital = 500 //first letter capital will make this public variable

func main() {
	// fmt.Println("variables")

	//uint8 can take value upto 255
	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("The type of variable is: %T \n", smallInt)

	//default values and aliases
	var someint int = 2556
	fmt.Println(someint)
	fmt.Printf("The type of variable is: %T \n", someint)

	//implicit type
	var somestring = "hey prerak"
	fmt.Println(somestring)
	// somestring = 3 //it'll give error since it already considers somestring as type=string

	//no var style
	//this is only allowed inside a function, not as global
	num := 34
	fmt.Println(num)

	fmt.Printf("The type of variable is: %T \n", FirstLetterCapital)
	fmt.Println(FirstLetterCapital)

}
