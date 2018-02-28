package main

import (
	"fmt"
)

type person struct {
	name    string
	surname string
	age     int
}

func main() {
	fmt.Println(person{"John", "Doe", 30})

	fmt.Println(person{name: "Jane", surname: "Doe", age: 40})

	fmt.Println(person{name: "Josh", surname: "Doe"})

	fmt.Println(&person{"Jack", "Doe", 50})

	p1 := &person{"Joel", "Doe", 50}
	fmt.Println("Person 1:", p1)

	p2 := person{"Jade", "Doe", 60}
	fmt.Println("Person 2:", p2)

	p2p := &p2
	fmt.Println("Person 2 (ptr):", p2p)

	p2.name = "Jake"
	p2.age = 61
	fmt.Println("Person 2:", p2)

	p2p.name = "Jean"
	p2p.age = 62
	fmt.Println("Person 2:", p2)
}
