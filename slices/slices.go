package main

import (
	"fmt"
)

func main() {
	var slice1 []string
	fmt.Println("slice1:", slice1)

	var slice2 = make([]string, 5)
	fmt.Println("slice2:", slice2)

	slice3 := make([]float32, 3)
	fmt.Println("slice3:", slice3)

	slice4 := []string{"a", "b", "c"}
	fmt.Println("slice4:", slice4)

	slice4[0] = "d"
	slice4[1] = "e"
	slice4[2] = "f"

	fmt.Println("slice4-set:", slice4)
	fmt.Println("slice4-get:", slice4[0])
	fmt.Println("slice4-len:", len(slice4))
	fmt.Println("slice4-cap:", cap(slice4))

	slice5 := make([]string, len(slice4))
	copy(slice5, slice4)
	fmt.Println("slice5:", slice5)

	slice5 = append(slice5, "g")
	slice5 = append(slice5, "h", "i")
	fmt.Println("slice5:", slice5)
	fmt.Println("slice5-len:", len(slice5))
	fmt.Println("slice5-cap:", cap(slice5))

	slice6 := slice5[1:5]
	fmt.Println("slice6:", slice6)
	fmt.Println("slice6-len:", len(slice6))
	fmt.Println("slice6-cap:", cap(slice6))

	slice7 := slice5[:3]
	fmt.Println("slice7:", slice7)
	fmt.Println("slice7-len:", len(slice7))
	fmt.Println("slice7-cap:", cap(slice7))

	slice8 := slice5[3:]
	fmt.Println("slice8:", slice8)
	fmt.Println("slice8-len:", len(slice8))
	fmt.Println("slice8-cap:", cap(slice8))

	slice9 := make([][]int, 3)
	fmt.Println("slice9:", slice9)
	for i := 0; i < len(slice9); i++ {
		len := len(slice9)
		slice9[i] = make([]int, len)
		for j := 0; j < len; j++ {
			slice9[i][j] = i + j
		}
	}
	fmt.Println("slice9:", slice9)
}
