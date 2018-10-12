package main

import (
	"fmt"
)

func main(){
	var i interface{} = "hello"

	/* statement asserts that the interface value i holds the concrete type - string.
	   and assigns the value of this string to variable s.
	 */
	s := i.(string)
	fmt.Println(s)


	/*
		To Test whether an interface value holds a specific type:

		a type assertion can return two values: the underlying value - s
		and a boolean value that reports whether the assertion succeeded - ok.
	 */
	s, ok := i.(string)
	fmt.Println(s, ok)


	/*
		If i holds a type - float64, then f will be the underlying value and ok = true.

		If not, ok  = false and f will be the zero value of type - float64 (in this case 0), and no panic occurs.
	 */
	f, ok := i.(float64)
	fmt.Println(f, ok)

	/*
	f = i.(float64) //panic

	due to not creating and assigning boolean value to check whether type assertion was successful.
	*/
}
