package main

import "fmt"

const (
	a = 20 + iota
	b = 20 + iota
	c = 20 + iota
	d = 20 + iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
