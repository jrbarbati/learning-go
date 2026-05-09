package main

import (
	"errors"
	"fmt"
	"strconv"
)

var operations = map[string]op{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

type op func(int, int) (int, error)

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "/", "0"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Printf("Invalid expression %v\n", expression)
			continue
		}

		p1, err := strconv.Atoi(expression[0])

		if err != nil {
			fmt.Println(err)
			continue
		}

		operation, ok := operations[expression[1]]

		if !ok {
			fmt.Printf("unsupported operator: %v\n", expression[1])
			continue
		}

		p2, err := strconv.Atoi(expression[2])

		if err != nil {
			fmt.Println(err)
			continue
		}

		result, err := operation(p1, p2)

		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(result)
	}
}

func add(a, b int) (int, error) {
	return a + b, nil
}

func sub(a, b int) (int, error) {
	return a - b, nil
}

func mul(a, b int) (int, error) {
	return a * b, nil
}

func div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}
