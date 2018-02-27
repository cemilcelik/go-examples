package main

import (
	"fmt"
)

func main() {
	map1 := make(map[string]int)
	fmt.Println("map1:", map1)

	map1["num1"] = 1
	map1["num2"] = 2
	fmt.Println("map1:", map1)

	map2 := map[string]int{"num1": 10, "num2": 20}
	fmt.Println("map2:", map2)

	num1 := map2["num1"]
	fmt.Println("num1:", num1)
	fmt.Println("len:", len(map2))

	delete(map2, "num1")
	fmt.Println("map2:", map2)
	fmt.Println("len:", len(map2))

	num2, prs := map2["num2"]
	fmt.Println("num2:", num2)
	fmt.Println("prs:", prs)
}
