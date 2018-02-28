package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "message 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "message 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("c1:", msg1)
		case msg2 := <-c2:
			fmt.Println("c2:", msg2)
		}
	}
}
