package main

import "fmt"

func main() {
	// declare a new variable a and assign it the value 200
	a := 200

	// declare a variable b and assign it (point it to) memory location of a
	/*
		**REMEMBER** | We don't know the exact memory location where a is stored
	 */
	b := &a


	// b contains the memory location of a but...
	// **KEY** | we want to increment the value STORED in a
	// TO DO this we must use *b - to dereference b. ALL this means is we follow the pointer from b to a,
	// 												add 1 to the value in a,
	//												and store this value back in the memory location stored in b- which is the memory location of a.
	*b++
	fmt.Print(a)
	
}
