package main

import "fmt"

// Closing a channel indicates that no more values will be sent on it.
// This can be useful to communicate completion to the channel’s receivers.

func main() {
	jobs := make(chan int, 5)
	// When we have no more jobs for the worker we’ll close the jobs channel.

	done := make(chan bool)

	// worker goroutine
	// It repeatedly receives from jobs with j, more := <-jobs.
	// In this special 2-value form of receive, the more value
	// WILL BE FALSE IF THE JOBS CHANNELS HAVE BEEN CLOSED
	// AND IF ALL VALUES IN THE CHANNEL HAVE ALREADY BEEN RECEIVED.
	// We use this to notify on done when we’ve worked all our jobs.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// This sends 3 jobs to the worker over the jobs channel, then closes it.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// We await the worker using the synchronization approach
	<-done
}
