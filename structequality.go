package main

import (
	"fmt"
)

// START OMIT
type Test struct {
	field1 string
	field2 int
}

func main() {
	x := Test{"foo", 1}
	y := Test{"foo", 1}

	fmt.Println(" Comparing structs:", x == y)

	a := &Test{"foo", 1}
	b := &Test{"foo", 1}

	fmt.Println("Comparing pointers:", a == b)
}

// END OMIT
