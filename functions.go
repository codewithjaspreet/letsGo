package main

import "fmt"

func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

// multiple return function

func vals() (int, int) {
	return 3, 7
}

func sums(a ...int) {

	fmt.Print(a, "")
	total := 0

	for b := range a {
		total += b
	}

	fmt.Print("total is ", total)
}
