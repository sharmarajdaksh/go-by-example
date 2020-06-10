package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// Dump a string (or just bytes) into a file.
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	// Create a new file / Open it for writing
	f, err := os.Create("/tmp/dat2")
	check(err)

	// Defer closing the file
	defer f.Close()

	// You can Write byte slices as you’d expect.
	d2 := []byte{115, 111, 109, 101, 10} // Translates to "some"
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// A WriteString is also available.
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Issue a Sync to flush writes to stable storage.
	f.Sync()

	// bufio provides buffered writers in addition to
	// buffered readers.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	// Use Flush to ensure all buffered operations have
	// been applied to the underlying writer.
	w.Flush()

}
