package main

import "fmt"

const s string = "constant"

func main() {
	fmt.Println("s =", s)

	const f float64 = 5.5
	fmt.Println("f =", f)

	const d = 3e18
	fmt.Println("d =", d)
	fmt.Println("d =", uint64(d))
}
