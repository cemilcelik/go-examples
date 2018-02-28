package main

import (
	"fmt"
	"time"
)

func f(c chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done...")

	c <- true
}

func main() {
	channel := make(chan bool, 1)
	go f(channel)

	r := <-channel

	fmt.Println("result:", r)
}
