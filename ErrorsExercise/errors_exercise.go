package main

import (
	"fmt"
	"math"
)
const e = 1e-8  // small delta

// created a new type
type ErrNegativeSqrt float64

// made it an error by using the Error() method from the standard: interface error { Error() string}
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f\n", e)
}


func Sqrt(x float64) (float64, error){

	if x < 0 {
		// this works because in go you specify the value of float64 by: float64(n)
		return 0, ErrNegativeSqrt(x)
	}
	// code that computes the square root
	z := x
	for {
		new_z := z - ((z*z - x) / (2*z))
		if math.Abs(new_z - z) < e {
			return new_z, nil
		}
		z = new_z
	}

}


func main(){
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
