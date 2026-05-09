package main

import "math/rand"

func main() {
	nums := make([]int, 0, 100)

	for i := 0; i < 100; i++ {
		nums = append(nums, rand.Intn(100))
	}
}
