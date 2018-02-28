package main

import (
	"fmt"
)

func f(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s, i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(s string) {
		fmt.Println(s)
	}("going")

	fmt.Scanln()
	fmt.Println("Done")
}
