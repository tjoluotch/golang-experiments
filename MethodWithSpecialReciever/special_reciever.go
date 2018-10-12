package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// method with special receiver function
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	// method abs takes the particular Vertex object
	fmt.Println(v.Abs())
}