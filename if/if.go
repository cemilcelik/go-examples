package main

import "fmt"

func main() {
	// if 7%2 == 0 {
	// 	fmt.Println("7 is even.")
	// } else {
	// 	fmt.Println("7 is odd.")
	// }

	if num := -5; num < 0 {
		fmt.Println(num, "is negative.")
	} else if num > 0 {
		fmt.Println(num, "is positive.")
	} else {
		fmt.Println(num, "is notr.")
	}
}
