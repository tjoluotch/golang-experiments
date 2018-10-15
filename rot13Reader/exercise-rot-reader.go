package main

import (
	"bytes"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

var ascii_uppercase = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var ascii_lowercase = []byte("abcdefghijklmnopqrstuvwxyz")
var ascii_uppercase_len = len(ascii_uppercase)
var ascii_lowercase_len = len(ascii_lowercase)

func rot13(b byte) byte {
	// returns the index of the 1st instance of b variable in ascii_uppercase
	// or -1 if b isn't present
	pos := bytes.IndexByte(ascii_uppercase, b)
	// e.g L makes pos = 11
	if pos != -1 {
		// 11+ 13 = 24
		// 24 % 26 = 24
		// ascii_uppecase[24] = Y
		return ascii_uppercase[(pos+13) % ascii_uppercase_len]
	}
	pos = bytes.IndexByte(ascii_lowercase, b)
	if pos != -1 {
		return ascii_lowercase[(pos+13) % ascii_lowercase_len]
	}
	return b
}

func(r *rot13Reader) Read(p []byte) (n int,err error){
	// read from to IO reader's implementation of Read
	n, err = r.r.Read(p)
	for i := 0; i < n; i++ {
		// replace existing byte with rot13 byte
		p[i] = rot13(p[i])
	}
	// return the number of bytes read and an error
	return n, err
}


func main(){
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	// copy Readers contents from r, and writes them to os.Stdout
	io.Copy(os.Stdout, &r)
}
