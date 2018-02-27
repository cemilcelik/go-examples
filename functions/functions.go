package main

import (
	"fmt"
)

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res1 := plus(3, 5)
	fmt.Println("3 + 5 =", res1)

	res2 := plusPlus(3, 5, 7)
	fmt.Println("3 + 5 + 7 =", res2)
}
