package main

import "fmt"

func main() {
	x := [5]int{2, 3, 4, 5, 6}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
