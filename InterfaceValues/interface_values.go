package main

import (
	"fmt"
	"math"
)

/*
an Interface can be thought of as a tuple of a value and a concrete type (value, type)

an Interface holds a value of a specific underlying concrete type
 */

type I interface {
	M()
}

type T struct {
	S string
}

type K struct {
	defn string
}

// pointer receiver of type T struct implements Interface I method M
func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

// Value receiver of type F float64 implements interface I method M
func(f F) M() {
	fmt.Println(f)
}

//pointer receiver of type K struct implements Interface I method M
func (g *K) M() {
	if g == nil{
		fmt.Println("this is nil")
		return
	}
	fmt.Println(g.defn)
}

func main()  {
	/*
	General notes on an Interfaces versatility:
	An interface is very versatile

	- i is declared as a variable of type I (which is an interface)

	- i is then assigned a value of pointer of type T struct
	- because of the assignment...
	- i implements method M using the pointer receiver type T struct

	- i is then assigned a value of type F float64
	- because of the assignment...
	- i implements method M using the value receiver type F

	- i is then assigned a value of pointer K struct
	- because of the assignment..
	- i implements method M using the pointer receiver type K struct
	 */


	var i I

	i = &T{"hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()

	var check *K
	i = check
	describe(i)
	i.M()

}

func describe(i I){
	fmt.Printf("(%v, %T)\n", i, i)
}
