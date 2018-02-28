package main

import (
	"fmt"
)

func main() {
	messageChan := make(chan string)

	go func() { messageChan <- "hello" }()

	msg := <-messageChan

	fmt.Println("msg:", msg)
}
