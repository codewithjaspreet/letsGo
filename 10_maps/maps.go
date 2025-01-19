package main

import (
	"fmt"
	"maps"
)

func main() {

	m := make(map[string]string)

	// setting an element

	m["name"] = "golang"
	m["area"] = "backend"

	// get an element

	fmt.Println(m["name"], m["area"])

	// imp : if key does not exist in map it return zero value
	// zeroed values -  string - "" , bool  - false  , int - 0

	n := make(map[string]int)

	fmt.Println(n["phone"]) // return 0

	fmt.Println(len(m))

	delete(n, "phone")

	// without make

	t1 := map[string]int{"price": 40, "phone": 3}
	t2 := map[string]int{"price": 40, "phone": 3}

	fmt.Println(t1)

	v, ok := t1["phone"] // useful for checking value and if value exist

	fmt.Println(v)

	if ok {
		fmt.Println("ALL OK")
	} else {
		fmt.Println("Not Ok")
	}

	fmt.Println(maps.Equal(t1, t2))
}
