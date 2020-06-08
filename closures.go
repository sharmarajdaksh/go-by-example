package main

import "fmt"

func intSeq() func() int {
	i := 0

	// The returned function closes
	// over the variable i to form a closure.
	return func() int {
		i++
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt()) // 1
	fmt.Println(nextInt()) // 2
	fmt.Println(nextInt()) // 3

	fmt.Println()

	newInts := intSeq()
	fmt.Println(newInts()) // 1
}
