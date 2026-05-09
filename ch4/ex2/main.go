package main

import (
	"fmt"
	"math/rand"
)

func main() {
	nums := make([]int, 0, 100)

	for i := 0; i < 100; i++ {
		nums = append(nums, rand.Intn(100))
	}

	for _, n := range nums {
		switch {
		case n%2 == 0 && n%3 == 0:
			fmt.Println("Six!")
		case n%2 == 0:
			fmt.Println("Two!")
		case n%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}
	}
}
