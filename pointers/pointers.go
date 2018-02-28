package main

import (
	"fmt"
)

func changeValue1(i int) {
	i = 0
}

func changeValue2(i *int) {
	*i = 0
}

func main() {
	data := 5
	fmt.Println("data:", data)

	changeValue1(data)
	fmt.Println("data:", data)

	changeValue2(&data)
	fmt.Println("data:", data)

	fmt.Println("addr:", &data)
}
