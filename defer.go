package main

import (
	"bufio"
	"fmt"
	"os"
)

// START OMIT
func CountWords(filename string) (int, error) {
	f, err := os.Open(filename)

	if err != nil {
		return 0, err
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	count := 0

	for scanner.Scan() {
		count++
	}

	return count, nil
}

// END OMIT

func main() {
	count, err := CountWords("somefile.txt")

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Found %d words\n", count)
	}
}
