package main

import "fmt"

func main() {
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}

	a := greetings[:2]
	b := greetings[1:4]
	c := greetings[3:]

	fmt.Println(greetings)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// BONUS
	// Be careful when using slices of slices or slices of arrays, as they share memory

	a[1] = "Bonjour"
	a = append(a, "Ciao")

	fmt.Println(greetings)
	fmt.Println(a)
}
