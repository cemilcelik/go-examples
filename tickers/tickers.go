package main

import (
	"fmt"
	"time"
)

func main() {
	ticker1 := time.NewTicker(300 * time.Millisecond)
	fmt.Printf("type: %T\n", ticker1.C)

	go func() {
		for t := range ticker1.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(1500 * time.Millisecond)
	ticker1.Stop()
	fmt.Println("ticker1 stopped.")
}
