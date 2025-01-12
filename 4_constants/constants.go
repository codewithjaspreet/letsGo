package main

import (
	"fmt"
	"math"
)

// this is allowed not , infer waay - X s:= "constant" X
const s string = "constant"

func main() {

	fmt.Println(s)
	const n = 5000
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))

}
