package main

import (
	"fmt"
	"time"
)

func worker(i int, jobs chan int, results chan int) {
	for j := range jobs {
		fmt.Printf("worker:%d, job:%d started.\n", i, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d, job:%d ended.\n", i, j)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 6)
	results := make(chan int, 6)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 1; i <= 5; i++ {
		<-results
	}
	close(results)
}
