package main

import "fmt"

// When using channels as function parameters,
// you can specify if a channel is meant to only send
// or receive values.
// This specificity increases the type-safety of the program.

// ping() can only SEND to the channel
// Receiving from the channel could cause
// a compile-time error
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pong() can only SEND to pongs and only
// RECEIVE from pings
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
	ping(pings, "second passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
