package main

import (
	"fmt"
	"time"
)

// We can use channels to synchronize execution across goroutines.

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second * 3)
	fmt.Println("done")

	done <- true
}

func main() {

	done := make(chan bool, 1)
	go worker(done)

	// The execution of the program will only complete
	// When a message is received on the channel from
	// the worker routine
	<-done
	// Without this line, program would exit
	// without the worker routine doing anything
}
