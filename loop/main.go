package main

import "fmt"

func main() {
	week := []string{"sun", "mon", "tue", "wed", "thu", "fri", "sat"}

	// for d := 0; d < len(week); d++ {
	// 	fmt.Println(week[d])
	// }

	for i, v := range week { //range returns index and value
		fmt.Println(week[i], v)
		if i == 3 {
			goto mail
		}
	}

mail:
	fmt.Println(("this is example of goto"))

}
