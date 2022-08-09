package main

import "fmt"

func main() {
	fmt.Println("welcome to fucntions")
	user := "prerak"
	greeter(user)
	res := adder(3, 4)
	fmt.Println("result:", res)
	prores, _ := proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("pro result is:", prores)

}

func greeter(user string) {
	fmt.Println("welcome", user)
}

func adder(a int, b int) int {
	return a + b
}

func proAdder(values ...int) (int, string) {
	total := 0
	for num := range values {
		total += num
	}
	return total, "string result"

}

// func (){
// 	fmt.printf("hey\n")
// }()
