package main

import "fmt"

func main() {
	var b byte = 255                             // Idiomatic way of assigning a literal to a byte, avoid b := byte(255)
	var smallerI int32 = 2_147_483_647           // Idiomatic way of assigning a literal to an int32, avoid smallerI := int32(2_147_483_647)
	var bigI uint64 = 18_446_744_073_709_551_615 // Idiomatic way of assigning a literal to an uint64, avoid bigI := uint64(18_446_744_073_709_551_615)

	fmt.Printf("%v %v %v\n", b, smallerI, bigI)
	fmt.Printf("%v %v %v\n", b+1, smallerI+1, bigI+1) // Overflow: 0 -2,147,483,648 0
}
