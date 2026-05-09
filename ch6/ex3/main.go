package main

import (
	"fmt"
	"time"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	{
		start := time.Now()

		persons := make([]*Person, 0)

		for range 10_000_000 {
			persons = append(persons, &Person{"Bob", "Maria", 30})
		}

		end := time.Now()

		fmt.Println(end.Sub(start))
	}
	{
		start := time.Now()

		persons := make([]*Person, 0, 10_000_000)

		for range 10_000_000 {
			persons = append(persons, &Person{"Bob", "Maria", 30})
		}

		end := time.Now()

		fmt.Println(end.Sub(start))
	}
	{
		start := time.Now()

		persons := make([]*Person, 10_000_000)

		for i := range 10_000_000 {
			persons[i] = &Person{"Bob", "Maria", 30}
		}

		end := time.Now()

		fmt.Println(end.Sub(start))
	}
	{
		start := time.Now()

		persons := make([]Person, 10_000_000)

		for i := range 10_000_000 {
			persons[i] = Person{"Bob", "Maria", 30}
		}

		end := time.Now()

		fmt.Println(end.Sub(start))
	}
}
