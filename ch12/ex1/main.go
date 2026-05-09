package main

import (
	"fmt"
	"sync"
)

func main() {
	numCh := make(chan int)
	wwg := sync.WaitGroup{}
	rwg := sync.WaitGroup{}

	wwg.Add(2)

	go func() {
		writeNums(1, 10, numCh)
		wwg.Done()
	}()

	go func() {
		writeNums(11, 20, numCh)
		wwg.Done()
	}()

	go func() {
		wwg.Wait()
		close(numCh)
	}()

	rwg.Add(1)
	go func() {
		readNums(numCh)
		rwg.Done()
	}()

	rwg.Wait()
}

func writeNums(start, finish int, ch chan<- int) {
	for i := start; i <= finish; i++ {
		ch <- i
	}
}

func readNums(ch <-chan int) {
	for num := range ch {
		fmt.Println(num)
	}
}
