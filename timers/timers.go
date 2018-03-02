package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(3 * time.Second)

	time1 := <-timer1.C
	fmt.Println("timer1 expired.", time1)

	timer2 := time.NewTimer(1 * time.Second)

	go func() {
		time2 := <-timer2.C
		fmt.Println("timer2 expired.", time2)
	}()

	stop := timer2.Stop()
	if stop {
		fmt.Println("timer2 stopped.", time.Now())
	}
}
