package main

// a function returing a function

func intSeq() func() int {

	i := 0

	return func() int {

		i++
		return i
	}

}
