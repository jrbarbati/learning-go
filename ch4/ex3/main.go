package main

import "fmt"

func main() {
	var total int

	for i := 0; i < 10; i++ {
		total := total + i
		fmt.Println(total) // This prints 0 - 9, total (line 9) is being set with := which makes a new total var each loop, shadowing the total var on line 6
		// So the total + i is using the total declared on line 6, which is always the zero value of an int (0), as it is never being updated.
	}

	fmt.Println(total)
}
