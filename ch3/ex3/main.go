package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	id        int
}

func main() {
	emp1 := Employee{"Joseph", "Barbati", 1}
	emp2 := Employee{firstName: "Isadora", lastName: "Blaschke da Silva", id: 2}

	var emp3 Employee
	emp3.firstName = "Maya"
	emp3.lastName = "Barbati"
	emp3.id = 3

	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp3)
}
