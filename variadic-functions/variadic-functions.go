package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Printf("%T %d ", nums, nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)
}

func main() {
	sum(1)
	sum(1, 2)
	sum(1, 2, 3)
	sum(1, 2, 3, 4)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)
}
