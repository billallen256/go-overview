package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"
)

// START OMIT
func main() {
	f, err := os.Open("somefile.txt")

	if err != nil {
		log.Fatal(err.Error())
	}

	defer f.Close()

	go func() {
		b, err := ioutil.ReadAll(f)

		if err != nil {
			log.Println(err.Error())
			return
		}

		for _, r := range bytes.Runes(b) {
			log.Println(r, string(r))
		}
	}()

	log.Println("Done")
}

// END OMIT
