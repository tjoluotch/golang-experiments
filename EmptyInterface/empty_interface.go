package main

import "fmt"

func main(){

// The interface that specifies zero methods is known as the empty interface e.g interface{}

/*
An empty interface can hold a value of ANY type

e.g fmt.Print() takes any number of arguments of type interface{}
 */
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}){

	//print value and type of interface argument
	fmt.Printf("(%v, %T)\n", i, i)
}
