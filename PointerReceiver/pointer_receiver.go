package main

/*
Reasons to use Pointer Receiver:
1) so that the method can modify the value that its receiver points to.
2) to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct,
for example.
*/


/*
In general, all methods on a given type should have either value or pointer receivers,
but NOT a mixture of both.
 */


import (
	"fmt"
	"math"
)

type Vertex struct {
	X,Y float64
}
// Find the the Hypotenuse: root(x^2 + y^2)
//This is a method with a value receiver
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Take the current vertex object and multiply X and Y values by the Scale Factor
// This method has a pointer receiver which means modify the exact vertex value in main
// The Scale method must have a pointer receiver to change the Vertex value declared in the main function.
func (v *Vertex) Scale(s_factor float64) {
	v.X = v.X * s_factor
	v.Y = v.Y * s_factor
}

// func v Vertex) Scale(s_factor float64) {} is an example of a Value Receiver
// With a value receiver, the Scale method operates on a copy of the original Vertex value.
func main () {
	v := Vertex{3,4}
	fmt.Println(v.Abs())
	v.Scale(10)
	fmt.Println(v.Abs())
}
