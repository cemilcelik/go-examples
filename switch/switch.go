package main

import (
	"fmt"
	"time"
)

func main() {
	num1 := 3
	fmt.Print("Write ", num1, " as ")
	switch num1 {
	case 1:
		fmt.Println("one.")
	case 2:
		fmt.Println("two.")
	case 3:
		fmt.Println("three.")
	}

	weekday := time.Now().Weekday()
	switch weekday {
	case time.Saturday, time.Sunday:
		fmt.Println("It's a weekend.")
	default:
		fmt.Println("It's a weekday.")
	}

	var t = time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon.")
	case t.Hour() > 12:
		fmt.Println("It's after noon.")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool.")
		case int:
			fmt.Println("I'm a int.")
		case string:
			fmt.Println("I'm a string.")
		default:
			fmt.Printf("Don't know type: %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI("test")
	whatAmI(10)
	whatAmI(5.5)

	fmt.Printf("whatAmI is a: %T", whatAmI)
}
