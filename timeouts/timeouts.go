package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "hello"
	}()

	select {
	case msg1 := <-c1:
		fmt.Println("msg1:", msg1)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout c1")
	}

	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "hi"
	}()

	select {
	case msg2 := <-c2:
		fmt.Println("msg2:", msg2)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 2")
	}
}
