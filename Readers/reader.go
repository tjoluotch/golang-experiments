package main

import (
	"fmt"
	"io"
	"strings"
)

/*
ASCII and Unicode:

ASCII can be thought of as a subset of unicode.

The Unicode standard team intentionally chose to assign the first 256 unicode characters to match Latin-1.
The first 128 characters of Latin-1 match standard ASCII.

This means that if UTF-8 encoding is used, the first 128 Unicode character are Identical to ASCII and the first 256
are identical to Latin-1.
 */


/*
type Reader interface {
  Read(p []byte) (n int, err error)
}

a Reader is any type that implements the `Read` method.


For those unfamiliar with these interfaces; you pass in a slice of bytes,
and Read is asked to fill it with its data — which is does until it runs out of data.

 It returns the number of bytes read (in `n`) or an error if something goes wrong.
Additionally, if it has finished reading, it will return a special marker error called `io.EOF` (end of file).
 */

func main(){

	// The example code creates a strings.Reader and consumes its output 8 bytes at a time.

	// created strings reader
	r := strings.NewReader("Hello Reader")

	// byte slice of length 8
	b := make([]byte, 8)

	for{
		//The full reader is 12 bytes - "Hello Reader" : one byte for each char including whitespace
		// bytes are read as ASCII e.g H = 72, e = 101 etc...

		// populates the given byte slice with data - maximum 8 bytes can be read at a time at a time
		n, err := r.Read(b)
		// First set of Bytes read are "Hello Re"
		fmt.Printf("n = %v  err = %v  b = %v\n", n, err, b)
		// Next set of Bytes read are "ader"
		fmt.Printf("b[:n] = %q\n", b[:n])

		// finished reading the
		if err == io.EOF {
			break
		}
	}
}
