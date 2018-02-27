package main

import "fmt"

func main() {

	// fmt.Println("i:", i)
	// 	for i := 0; i < 5; i++ {
	// }

	// j := 0
	// for j < 10 {
	// 	fmt.Println("j:", j)
	// 	j++
	// }

	// k := 0
	// for {
	// 	fmt.Println("k:", k)
	// 	k++
	// 	if k >= 5 {
	// 		break
	// 	}
	// }

	for l := 0; l < 10; l++ {
		if l%2 == 0 {
			continue
		}
		fmt.Println("l:", l)
	}
}
