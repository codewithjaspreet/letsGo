package main

import (
	"fmt"
	"slices"
)

func slicesImpl() {

	// array example

	var s []string

	fmt.Println("intital", s, s == nil, len(s) == 0)

	// to make a empty slice use the make function

	s = make([]string, 3)

	fmt.Println("emp", s, "len", len(s), "cap", cap(s))

	// slices have append functionality

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("append slice", s)

	// copy slice source to destination using builtiin copy method
	c := make([]string, len(s))

	copy(c, s)

	t := []string{"g", "h", "i"}
	t2 := []string{"g", "h", "i"}

	// methods on slices
	if slices.Equal(t, t2) {

		fmt.Println("t==t2")
	}

}
