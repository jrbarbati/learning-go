package main

import "fmt"

type Doubler interface {
	~int | ~int8 | ~uint8 | ~int16 | ~uint16 | ~int32 | ~uint32 | ~int64 | ~uint64 | ~float32 | ~float64
}

func double[T Doubler](val T) T {
	return val * 2
}

func main() {
	var i int
	var f float64

	i = 5
	f = 18.3

	fmt.Printf("double(i)=%d\n", double(i))
	fmt.Printf("double(f)=%.2f\n", double(f))
}
