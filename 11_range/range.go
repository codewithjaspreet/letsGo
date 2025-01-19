package main

import "fmt"

func main() {

	// use of range keyword on slices -

	slice := []int{6, 7, 8}

	for i, num := range slice {
		fmt.Println("Element at index", i, "is", num)

	}

	// use of range keyword on maps -

	m1 := make(map[string]int)

	m1["a1"] = 1
	m1["a2"] = 2
	m1["a3"] = 3

	for k, v := range m1 {

		fmt.Println("Key -", k, "has value -", v)
	}

	// use of range keyword on maps -

	for i, c := range "Jaspreet" {

		fmt.Println(i, c)
	}

}
