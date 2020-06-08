package main

import (
	"fmt"
	"sync"
	"time"
)

// To wait for multiple goroutines to finish, we can use a wait group.

// NOTE: A WaitGroup must be passed to functions by pointer.
func worker(id int, wg *sync.WaitGroup) {

	// On return, notify the WaitGroup that we’re done.
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	// This WaitGroup is used to wait for all the goroutines launched here to finish.
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// Launch several goroutines and increment the WaitGroup counter for each.
		wg.Add(1)
		go worker(i, &wg)
	}

	// Block until the WaitGroup counter goes back to 0;
	// all the workers notified they’re done.
	wg.Wait()
}
