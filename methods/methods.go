package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

func (r rect) area() int {
	return r.width * r.height
}

func (r *rect) perim() int {
	return r.width*2 + r.height*2
}

func main() {
	rect1 := rect{5, 10}

	area1 := rect1.area()
	fmt.Println("area1:", area1)

	rect2 := &rect1

	rect2.width = 2
	rect2.height = 4

	area2 := rect2.area()
	fmt.Println("area2:", area2)

	fmt.Println("rect1:", rect1)
	fmt.Println("rect2:", rect2)

	rect3 := rect{width: 4, height: 8}
	fmt.Println("rect3:", rect3)
	rect3p := &rect3
	area3 := rect3p.area()
	perim3 := rect3p.perim()
	fmt.Println("area3:", area3)
	fmt.Println("perim3:", perim3)
}
