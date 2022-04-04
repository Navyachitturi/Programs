package main

import "fmt"

// we create values of a certain type that are stored in variables
// and those variables have identifiers
var x int

type person struct {
	first string
	last  string
}

//type foo int
//var y foo
// const bar = 42
// y = 42
// fmt.printf("%T", int(y))
// fmt.printf("%T\n", bar)
// fmt.printf(bar)

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
	}
	fmt.Println(p1)
}
