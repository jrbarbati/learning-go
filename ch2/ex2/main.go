package main

import "fmt"

func main() {
	const value = 45      // Idiomatic way of initializing a const as this gives the most flexibility
	var i = value         // Idiomatic way of assigning an int to the value of a const as the default type of the literal is int
	var f float64 = value // Idiomatic way of assigning a float the value of the const as the default of the literal is int

	fmt.Printf("%d %f\n", i, f)
}
