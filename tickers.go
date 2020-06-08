package main

import (
	"fmt"
	"time"
)

// Timers are for when you want to do something once in the future.
// Tickers are for when you want to do something repeatedly at regular intervals

func main() {

	// Tickers use a similar mechanism to timers: a channel that is sent values.
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	// Tickers can be stopped like timers. Once a ticker is
	// stopped it won’t receive any more values on its channel.
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
