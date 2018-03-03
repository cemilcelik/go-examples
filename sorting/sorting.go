package main

import (
	"fmt"
	"sort"
)

func main() {
	var1 := []int{5, 10, 3, 15}
	sort.Ints(var1)
	fmt.Println("var1:", var1)

	isIntsAreSorted := sort.IntsAreSorted(var1)
	fmt.Println("isIntsAreSorted:", isIntsAreSorted)

	i := sort.SearchInts(var1, 10)
	fmt.Println("search i:", i)

	var2 := []string{"bb", "01", "a", "aaa", "1", "b", "aa", "10", "bbb", "2", "9", "11", "19", "20", "21"}
	sort.Strings(var2)
	fmt.Println("var2:", var2)
}
