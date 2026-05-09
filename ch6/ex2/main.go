package main

import "fmt"

func UpdateSlice(slice []string, str string) {
	slice[len(slice)-1] = str
	fmt.Println("Inside UpdateSlice(): ", slice)
}

func GrowSlice(slice []string, str string) {
	slice = append(slice, str)
	fmt.Println("Inside GrowSlice(): ", slice)
}

func main() {
	s := []string{"a", "b", "c"}
	UpdateSlice(s, "d")
	fmt.Println("in main after UpdateSlice:", s)
	GrowSlice(s, "e")
	fmt.Println("in main, after GrowSlice:", s)
}
