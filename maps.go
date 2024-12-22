package main

import "fmt"

func mapsImplementation() {

	// initialise map
	m := make(map[string]int)

	// assign key value pairs
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// if not exist then it will be 0 value by default
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	delete(m, "k2")
	fmt.Println("map", m)

	clear(m)

	fmt.Println("map", m)

}
