package main

import "fmt"

func main() {
	message := "Hello 👩 and 👨"

	runes := []rune(message)

	fmt.Println(string(runes[3]))

	// Bonus
	// Bytes and runes are different with non-UTF-8 chars

	bytes := []byte(message)

	fmt.Printf("Runes (len(runes) = %v) %v\n", len(runes), runes) // a rune holds enough space (4 bytes?) to hold a UTF-16 or UTF-32 char
	fmt.Printf("Bytes (len(bytes) = %v) %v\n", len(bytes), bytes) // a byte only holds 8 bits so the larger chars are spread across multiple bytes

	// slices of strings -- substring

	// Substrings like this are a little dangerous as you are substring-ing by bytes, not indexes

	s := "this is a string"
	slice := s[6:10]

	t := "Hello 👩 and 👨"
	tSlice := t[6:10]

	fmt.Println(slice)
	fmt.Println(tSlice)
}
