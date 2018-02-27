package main

import "fmt"

func main() {
	var a int
	a = 5
	fmt.Println("a =", a)

	var b string = "initial"
	fmt.Println("b =", b)

	var c, d int = 2, 4
	fmt.Printf("c = %d, d = %d \n", c, d)

	var e int // initial 0
	fmt.Println("e =", e)

	f := "short"
	fmt.Println("f =", f)

	g := true
	fmt.Println("g =", g)
}
