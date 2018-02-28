package main

import (
	"fmt"
)

func main() {
	bufferingChan := make(chan string, 2)

	bufferingChan <- "buffering"
	bufferingChan <- "channel"

	fmt.Println(<-bufferingChan)
	fmt.Println(<-bufferingChan)
}
