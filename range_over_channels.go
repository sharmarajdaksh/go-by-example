package main

import "fmt"

func main() {

	queue := make(chan string, 5)
	queue <- "one"
	queue <- "two"
	queue <- "3"
	queue <- "4"
	queue <- "5"
	close(queue)

	// Iterate over elements in the channel
	for elem := range queue {
		fmt.Println(elem)
	}
}
