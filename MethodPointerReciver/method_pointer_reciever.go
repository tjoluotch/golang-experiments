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

func(v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func(v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3,4}
	fmt.Printf("Before Scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After Scaling: %v, Abs: %v\n", v, v.Abs())
}
