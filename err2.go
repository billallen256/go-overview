package main

import (
	"errors"
	"fmt"
	"os"
)

// START MAIN OMIT
func wrap(e error) error {
	return fmt.Errorf("It was broke because of '%w'", e)
}

func main() {
	e := wrap(os.ErrPermission)
	for e != nil {
		fmt.Println(e)
		e = errors.Unwrap(e)
	}

	e = wrap(os.ErrNotExist)
	fmt.Println(errors.Is(e, os.ErrNotExist))

	e = wrap(&os.PathError{Op: "read", Path: "blah", Err: os.ErrNotExist})
	fmt.Println(e)
	pe := &os.PathError{}
	fmt.Println(errors.As(e, &pe))
	fmt.Println(pe.Error())
}

// END MAIN OMIT
