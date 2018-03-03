package main

import (
	"fmt"
	"os"
)

func createFile(p string) *os.File {
	fmt.Println("File creating..")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("File writing..")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("File closing..")
	f.Close()
}

func main() {
	f := createFile("./file.txt")
	defer closeFile(f)
	writeFile(f)
}
