package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {
	str := "Lorem ipsum dolor sit amet."

	p("str:", str)
	p("---")
	p("begins with 'L':", s.HasPrefix(str, "L"))
	p("contains 'm':", s.Contains(str, "m"))
	p("ends with '.':", s.HasSuffix(str, "."))
	p("index of 'm':", s.Index(str, "m"))
	p("join:", s.Join([]string{"foo", "bar"}, "-"))
	p("split:", s.Split(str, " "))
	p("count of 'm':", s.Count(str, "m"))
	p("upper:", s.ToUpper(str))
	p("lower:", s.ToLower(str))
	p("replace:", s.Replace(str, " ", "_", -1))
	p("replace:", s.Replace(str, " ", "_", 1))
	p("---")
	p("len:", len(str))
	p("str[1]:", str[1])
}
