package main

import (
	"fmt"
)

func main() {
	var arrA [5]int
	fmt.Println("arrA:", arrA)

	arrA[4] = 100
	fmt.Println("set: ", arrA)
	fmt.Println("get: ", arrA[4])
	fmt.Println("len: ", len(arrA))

	arrB := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arrB:", arrB)

	var arrC [5][5]int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			arrC[i][j] = i + j
		}
	}
	fmt.Println("arrC:", arrC)
}
