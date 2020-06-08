package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines.
func main() {

	// This channel will support string type values
    messages := make(chan string) 

	// Send message to channel from separate thread 
    go func() { messages <- "ping" }()

	// Get exactly ONE message from the channel
	// Will hold program execution until that message is received 
    msg := <-messages
    fmt.Println(msg)
}