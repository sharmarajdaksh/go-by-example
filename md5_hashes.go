package main

import (
	// Go implements several hash functions in various crypto/* packages.
	"crypto/md5"
	"fmt"
)

func main() {
	s := "sha1 this string"

	// The pattern for generating a hash is md5.New(),
	// md5.Write(bytes), then md5.Sum([]byte{}).

	h := md5.New()

	// Write expects bytes
	h.Write([]byte(s))

	// This gets the finalized hash result as a byte slice.
	// The argument to Sum can be used to append to an existing
	// byte slice: it usually isn’t needed.
	bs := h.Sum(nil)

	fmt.Println(s)

	// Hashed values are often printed in hex, for example in git commits.
	// Use the %x format verb to convert a hash results to a hex string.
	fmt.Printf("%x\n", bs)
}

// Running the program computes the hash and prints it in a human-readable hex format.
