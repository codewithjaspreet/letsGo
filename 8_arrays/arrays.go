package main

import "fmt"

func arrays() {

	// sample array initialisations

	var a [5]int
	fmt.Println("Temp:", a)

	a[4] = 100

	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// array initialise with size

	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println("dcl:", b)

	//[...] automatically define the size of the array
	d := [...]int{1, 2, 3, 4, 5}

	fmt.Println("dcl:", d)

}
