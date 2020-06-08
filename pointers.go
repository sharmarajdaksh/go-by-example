package main

import "fmt"

func incbyval(i int) {
	i++
	fmt.Println(i)
}

func incbyptr(i *int) {
	*i++
	fmt.Println(*i)
}

func main() {
	i := 5

	incbyptr(&i)   // 6
	fmt.Println(i) // 6

	incbyval(i)    // 7
	fmt.Println(i) // 6
}
