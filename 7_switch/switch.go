package main

import (
	"fmt"
	"time"
)

func main() {

	// simple switch use case

	i := 2

	fmt.Print("Write", i, " as ")

	switch i {

	case 1:
		fmt.Println("one")

	case 2:
		fmt.Println("two")

	case 3:
		fmt.Println("three")
	}

	// default keyword + multiple expression in same line in case

	switch time.Now().Weekday() {

	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	whatAmI := func(i interface{}) {

		switch t := i.(type) {

		case bool:
			fmt.Println("I am a bool")

		case int:
			fmt.Println("I am a bool")

		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
