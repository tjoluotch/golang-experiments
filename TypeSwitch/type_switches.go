package main

import "fmt"

func do(i interface{}) {
	// expression in type switch is evaluating the interface type
	switch v := i.(type) {
	// cases are testing different types
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about Type %T!\n", v)
	}
}

func main(){
	do(64)
	do("whats up")
	do(true)
}
