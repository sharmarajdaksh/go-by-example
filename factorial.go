package main

import "fmt"

func fact(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	m := 5
	fmt.Println(fact(m))

	n := 10
	fmt.Println(fact(n))

}
