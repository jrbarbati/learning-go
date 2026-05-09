package main

import (
	"fmt"
	"strconv"
)

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type PrintableIntF int

func (pi PrintableIntF) String() string {
	// 50x slower than strconv.Itoa() -- see main_test.go
	return fmt.Sprintf("%d", pi)
}

type PrintableInt int

func (pi PrintableInt) String() string {
	return strconv.Itoa(int(pi))
}

type PrintableFloat64 float64

func (pi PrintableFloat64) String() string {
	return fmt.Sprintf("%.3f", pi)
}

func Print[T Printable](p T) {
	fmt.Println(p)
}

func main() {
	var pi PrintableIntF
	var pf PrintableFloat64

	pi = 60
	pf = 183.1226

	Print(pi)
	Print(pf)
}
