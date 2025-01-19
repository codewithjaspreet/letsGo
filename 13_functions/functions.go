package main

import "fmt"

// Simple addition function
func add(a, b int) int {
	return a + b
}

// Function returning multiple values
func getLangs() (string, string, string) {
	return "golang", "java", "python"
}

// Function returning another function
func fnWithfnReturn() func(a int) int {
	return func(a int) int {
		return a * 2 // Example logic
	}
}

// Function accepting another function as an argument
func fnWithfnArg(fn func(a int) int) int {
	return fn(1)
}

func main() {
	// Calling the addition function
	res := add(3, 4)

	// Calling the function returning multiple values
	lang1, lang2, lang3 := getLangs()

	// Printing the results
	fmt.Println(lang1, lang2, lang3)
	fmt.Println(res)

	// Calling fnWithfnArg with an inline function
	fnres := fnWithfnArg(
		func(a int) int {
			return a + 5 // Example inline logic
		},
	)

	fmt.Println(fnres)

	// Example of calling fnWithfnReturn
	fn := fnWithfnReturn()
	fmt.Println(fn(3)) // Call the returned function
}
