package main

import (
	"fmt"
)

func main() {
	// START OMIT
	s := make([]string, 0, 2)
	s = append(s, "foo") // append a string to the slice at postion 0
	s = append(s, "bar") // append at position 1
	s = append(s, "baz") // grow and append at position 2

	t := s[1:] // t refers to the last to items in s

	fmt.Printf("s = %+v\n", s)
	fmt.Printf("t = %+v\n", t)

	t[0] = "boo"
	fmt.Printf("s = %+v\n", s)
	fmt.Printf("t = %+v\n", t)

	t = append(t, "who") // lengthens t slice but not s slice
	fmt.Printf("s = %+v\n", s)
	fmt.Printf("t = %+v\n", t)

	t[0] = "too"
	t[1] = "too"
	t[2] = "too"
	fmt.Printf("s = %+v\n", s)
	fmt.Printf("t = %+v\n", t)
	// END OMIT
}
