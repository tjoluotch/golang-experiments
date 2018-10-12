package main

import "fmt"

type Person struct {
	Name string
	Age int
}

// method that uses a value receiver of type Person
/*
	type Stringer interface{
		String() string
 */

 // Type that can describe itself as a string
 // changes the type of print message

//func (p Person) String() string {
//	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
//}


func(p *Person) typeDef() {
	fmt.Printf("I am type of %+v\n", p)
}

func main(){
	a := Person{"Aaron Judge", 29}
	w := Person{"Willie Coly Stein", 54}

	a.typeDef()
	w.typeDef()



	fmt.Println(a,w)
}


