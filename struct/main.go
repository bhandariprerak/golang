package main

import "fmt"

func main() {
	fmt.Println("welcome to structs in golang")
	prerak := User{"Prerak", "prerak@kaleyra.com", true, 24}
	fmt.Println(prerak)
	fmt.Printf("user details for prerak are: %+v\n", prerak)
	fmt.Printf("name is: %v and email is %v\n", prerak.Name, prerak.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
