package main

import "fmt"

// START MAIN OMIT
func main() {
	var i interface{} = "hello" // interface{} is the empty interface in go

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

// END MAIN OMIT
