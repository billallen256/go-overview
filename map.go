package main

import (
	"log"
)

func main() {
	// START OMIT
	m := make(map[string]int64)
	m["bar"] = 101

	for _, key := range []string{"foo", "bar"} {
		if value, exists := m[key]; exists {
			log.Printf("Found value for %s: %d", key, value)
		} else {
			log.Printf("Could not find %s", key)
		}
	}
	// END OMIT
}
