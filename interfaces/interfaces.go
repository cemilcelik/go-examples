package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circ struct {
	radius float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}

func (r *rect) perim() float64 {
	return r.width*2 + r.height*2
}

func (c *circ) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c *circ) perim() float64 {
	return c.radius * math.Pi * 2
}

func measure(g geometry) {
	fmt.Println("---")
	fmt.Printf("measure() type  : %T\n", g)
	fmt.Printf("measure() value : %v\n", g)
	fmt.Println("measure() area  :", g.area())
	fmt.Println("measure() perim :", g.perim())
}

func main() {
	rect1 := rect{width: 5, height: 10}
	fmt.Println("rect1:", rect1)
	fmt.Println("rect1 area:", rect1.area())
	fmt.Println("rect1 perim:", rect1.perim())

	circ1 := circ{radius: 4}
	fmt.Println("circ1:", circ1)
	fmt.Println("circ1 area:", circ1.area())
	fmt.Println("circ1 perim:", circ1.perim())

	measure(&rect1)
	measure(&circ1)
}
