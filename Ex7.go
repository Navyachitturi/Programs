package main

import "fmt"

func main() {
	x := "James Bond"
	if x == "Money" {
		fmt.Println(x)
	} else if x == "James Bond" {
		fmt.Println("BONDONBONDONBOND", x)
	} else {
		fmt.Println("neither")
	}
}
