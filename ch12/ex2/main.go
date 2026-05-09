package main

import (
	"fmt"
)

func main() {
	process()
}

func process() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		for i := 1; i <= 10; i++ {
			ch1 <- i
		}
	}()

	go func() {
		defer close(ch2)
		for i := 11; i <= 20; i++ {
			ch2 <- i
		}
	}()

	for ch1 != nil && ch2 != nil {
		select {
		case num1, ok := <-ch1:
			if !ok {
				ch1 = nil // Effectively turns off this select case
				continue
			}
			fmt.Println(num1)
		case num2, ok := <-ch2:
			if !ok {
				ch2 = nil // Effectively turns off this select case
				continue
			}
			fmt.Println(num2)
		}
	}
}
