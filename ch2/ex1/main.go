package main

import "fmt"

func main() {
	var i = 20 // Idiomatic way of setting a literal to an int as the default type of the literal is an int
	var f = float64(i)

	fmt.Println(i)
	fmt.Println(f)
}
