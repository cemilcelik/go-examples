package main

import (
	"fmt"
)

func vals() (int, int) {
	return 3, 5
}

func main() {
	v1, v2 := vals()
	fmt.Println("v1:", v1)
	fmt.Println("v2:", v2)

	ret1, _ := vals()
	fmt.Println("ret1:", ret1)

	_, ret2 := vals()
	fmt.Println("ret2:", ret2)
}
