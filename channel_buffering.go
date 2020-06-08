package main

import "fmt"

// By default channels are unbuffered,
// meaning that they will only accept sends (chan <-)
// if there is a corresponding receive (<- chan) ready
// to receive the sent value.

// Buffered channels accept a limited number of values
// without a corresponding receiver for those values.

//

func main() {
	messages := make(chan string, 2)

	// Because this channel is buffered,
	// we can send these values into the channel
	// without a corresponding concurrent receive.

	messages <- "buffered"
	messages <- "channel"
	// messages <- "THIS WILL CREATE AN ERROR"

	fmt.Println(<-messages)
	fmt.Println(<-messages) // You could comment this and the program would still work
	// fmt.Println(<-messages) // THIS ALSO WILL CREATE AN ERROR IF UNCOMMENTED
}
