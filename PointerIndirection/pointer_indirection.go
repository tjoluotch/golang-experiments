package main

import (
	"fmt"
	"math"
)

//Keys: Receivers can take either value or a pointer when called
//		Pointer Arguments must take pointer
// 		Value Arguments must take a value


type Vertex struct {
	X,Y float64
}

// Pointer receiver - can take either a value or a pointer as the receiver when called
// Scale x,y coordinates in vertex
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Pointer argument - must take a pointer
// Scale x,y coordinates in vertex
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}



// Function that has a Value Receiver - can take either a value or a pointer as the receiver when called
// abs function
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// function that takes value argument - must take a value of that specific type
//abs function
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	//whereas method with pointer receiver can take either a value or pointer receiver when called
	v.Scale(2)
	// function with a Pointer argument must take a pointer
	ScaleFunc(&v, 10)


	// function with a Pointer argument must take a pointer
	p := &Vertex{3,4}

	//whereas method with pointer receiver can take either a value or a pointer as the receiver when called
	p.Scale(3)
	ScaleFunc(p,8)

	fmt.Println(v, p)

	t := Vertex{3,4}
	fmt.Println(t.Abs())
	fmt.Println(AbsFunc(t))

	// pointer to vertex
	d := &Vertex{4,3}
	fmt.Println(d.Abs())
	fmt.Println(AbsFunc(*d))
}
