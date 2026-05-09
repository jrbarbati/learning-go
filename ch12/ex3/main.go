package main

import (
	"fmt"
	"math"
	"sync"
)

const (
	size = 100_000
)

var (
	sqrts = sync.OnceValue(buildMap)
)

func main() {
	for i := 0; i < size; i += 1000 {
		fmt.Println("sqrt of", i, "->", sqrts()[i])
	}
}

func buildMap() map[int]float64 {
	m := make(map[int]float64, size)

	for i := 0; i < size; i++ {
		m[i] = math.Sqrt(float64(i))
	}

	return m
}
