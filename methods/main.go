package main

import "fmt"

func main() {
	fmt.Println("welcome to structs in golang")
	prerak := User{"Prerak", "prerak@kaleyra.com", true, 24}
	fmt.Println(prerak)
	fmt.Printf("user details for prerak are: %+v\n", prerak) // %+v will give struct in a better way i.e, along with variables
	fmt.Printf("name is: %v and email is %v\n", prerak.Name, prerak.Email)
	prerak.GetStatus()
	prerak.ChangeMail()
	fmt.Printf("user details for prerak are: %+v\n", prerak) // %+v will give struct in a better way i.e, along with variables

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() { // only a copy of user is passed, not the actual value
	fmt.Println("user stautus is:", u.Status)

}

func (u *User) ChangeMail() { //since we used pointer here, now the original value has changed
	newMail := "pb@kaleyra.com"
	u.Email = newMail
}
