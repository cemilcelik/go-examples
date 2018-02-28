package main

import (
	"errors"
	"fmt"
)

type argErr struct {
	arg int
	txt string
}

func (e *argErr) Error() string {
	return fmt.Sprintf("%s - %d", e.txt, e.arg)
}

func newArgErr(arg int, txt string) *argErr {
	return &argErr{arg, txt}
}

func f1(num int) (int, error) {
	if num == 42 {
		return -1, errors.New("Can't work with 42.")
	}

	return num, nil
}

func f2(num int) (int, error) {
	if num == 42 {
		return -1, newArgErr(num, "Can't work with it.")
	}

	return num, nil
}

func main() {
	for _, v := range []int{7, 42} {
		fmt.Println(v)

		res, err := f1(v)
		fmt.Println("res:", res)
		fmt.Println("err:", err)
	}

	for _, v := range []int{7, 42} {
		fmt.Println(v)

		res, err := f2(v)
		fmt.Println("res:", res)
		fmt.Println("err:", err)
	}
}
