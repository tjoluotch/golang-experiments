package main

//A nil error denotes success; a non-nil error denotes failure.

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// pointer receiver - dereferences pointer of type MyError : in this case reads the value in the MyError pointer
/*
The error type is a built-in interface similar to fmt.Stringer:

type error interface {
    Error() string
}
 */
func (e *MyError) Error() string{
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error{
	return &MyError{
		time.Now(),
		"It Didn't work",
	}
}

func main(){
	if err := run(); err != nil {
		fmt.Println(err)
	}

	/*
	Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.

i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
	 */
}
