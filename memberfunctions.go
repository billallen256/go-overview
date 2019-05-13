package main

import (
	"fmt"
)

// START OMIT
type intSlice []int

func (s intSlice) FirstEvenItem() (int, error) {
	for _, x := range s {
		if x%2 == 0 {
			return x, nil
		}
	}

	return 0, fmt.Errorf("Could not find an even item")
}

func main() {
	s := intSlice{1, 3, 5, 6, 8, 10}
	firstEven, err := s.FirstEvenItem()

	if err == nil {
		fmt.Println(firstEven)
	}
}

// END OMIT
