package main

import "fmt"

func forLoops() {

	// this is simple for loop with a condition

	i := 1

	for i <= 3 {

		fmt.Println(i)

		i = i + 1
	}

	// direct initialisation within loop
	for j := 0; j < 3; j++ {

		fmt.Println(j)
	}

	// infinite loop , & break using break keyword

	for {
		fmt.Println("loop")
		break
	}

	// use of range keyword  range 6 - 0,1,2,3,4,5
	for n := range 6 {

		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}

}
