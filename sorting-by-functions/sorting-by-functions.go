package main

import (
	"fmt"
	"sort"
)

type sortByLength []string

func (s sortByLength) Len() int {
	return len(s)
}

func (s sortByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"banana", "apple", "kiwi", "cucumber"}
	fmt.Println("fruits:", fruits)

	sort.Sort(sortByLength(fruits))
	fmt.Println("fruits:", fruits)
}
