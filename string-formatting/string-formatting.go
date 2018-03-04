package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

var p = fmt.Printf

func main() {
	po := point{1, 2}

	p("%v \n", po)                     // {1 2}
	p("%+v \n", po)                    // {x:1 y:2}
	p("%#v \n", po)                    // main.point{x:1, y:2}
	p("%T \n", po)                     // main.point
	p("%t \n", true)                   // true
	p("%d \n", 123)                    // 123
	p("%b \n", 14)                     // 1110
	p("%c \n", 33)                     // !
	p("%x \n", 456)                    // 1c8
	p("%f \n", 78.9)                   // 78.900000
	p("%e \n", 123400000.0)            // 1.234000e+08
	p("%E \n", 123400000.0)            // 1.234000E+08
	p("%s \n", "\"string\"")           // "string"
	p("%q \n", "\"string\"")           // "\"string\""
	p("%x \n", "hex this")             // 6865782074686973
	p("%p \n", &p)                     // 0x533cc8
	p("|%6d|%6d| \n", 12, 345)         // |    12|   345|
	p("|%6.2f|%6.2f| \n", 1.2, 3.45)   // |  1.20|  3.45|
	p("|%-6.2f|%-6.2f| \n", 1.2, 3.45) // |1.20  |3.45  |
	p("|%6s|%6s| \n", "foo", "b")      // |   foo|     b|
	p("|%-6s|%-6s| \n", "foo", "b")    // |foo   |b     |

	s := fmt.Sprintf("a %s", "string")
	fmt.Println(s)                              // a string
	fmt.Fprintf(os.Stderr, "an %s \n", "error") // an error
}
