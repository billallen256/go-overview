package main

import (
	"log"
	"sync"
	"time"
)

// START MAIN OMIT
func main() {
	// Create a WaitGroup thats waiting for three tasks
	wg := &sync.WaitGroup{}
	wg.Add(3) // What happens if we change this?

	log.Println("Start")

	for i := 1; i < 4; i++ {
		// Launch three tasks as goroutines
		// This syntax is executing an anonymous function
		go func(amount int) {
			// Sleep the required number of seconds and then signal we are done
			seconds := time.Second * time.Duration(amount)
			time.Sleep(seconds)
			log.Println("Hello after", seconds)
			wg.Done()
		}(i)
	}

	// Code will block here until wg.Done() has been called 3 times
	wg.Wait() // What happens if we comment this?
	log.Println("Done")
}

// END MAIN OMIT
