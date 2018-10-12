package main

import "fmt"

/*
An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.
 */

 /*
 A type implements an interface by implementing its methods. There is no explicit declaration of intent,
 no "implements" keyword.

 Implicit interfaces decouple the definition of an interface from its implementation,
 which could then appear in any package without prearrangement.
  */

 type I interface {
 	M()
 }

 type Tiger struct {
 	S string
 }


// This method means type Tiger implements the interface I,
// but we don't need to explicitly declare that it does so.
 func (t Tiger) M(){
	 // print the String that is defined in type Tiger
	fmt.Println(t.S)
 }

 func main() {
 	var i I = Tiger{"Tony"}
	i.M()
 	i = Tiger{"Rex"}
 	i.M()

	 var next_one I = Tiger{"Stripey"}
	 next_one.M()
 }




