package main

import (
	"fmt"
)

func main() {
	slice1 := []string{"item1", "item2", "item3"}
	for _, v := range slice1 {
		fmt.Printf("value: %s\n", v)
	}

	map1 := map[string]string{"name": "John", "surname": "Doe"}
	for key, val := range map1 {
		fmt.Printf("%s: %s\n", key, val)
	}

	map2 := map[string]int{"k1": 1, "k2": 2}
	fmt.Println("map2:", map2)
	for k := range map2 {
		fmt.Printf("key: %s\n", k)
	}

	for index, unicode := range "golang" {
		fmt.Printf("index: %d, unicode: %d\n", index, unicode)
	}
}
