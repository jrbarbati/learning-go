package main

import (
	"fmt"
	"os"
)

func main() {
	size, err := fileLen("ch5/ex1/main.go")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(size)
}

func fileLen(filename string) (int, error) {
	file, err := os.Open(filename)

	if err != nil {
		return 0, err
	}
	defer file.Close()

	fi, err := file.Stat()

	if err != nil {
		return 0, err
	}

	return int(fi.Size()), nil
}
