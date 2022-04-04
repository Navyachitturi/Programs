package main

import "fmt"

type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	p := person{
		first: "James",
		last:  "Bond",
		favFlavors: []string{
			"chocklate",
			"rum and coke",
		},
	}
	fmt.Println(p.first)
	fmt.Println(p.last)
	for i, v := range p.favFlavors {
		fmt.Println(i, v)
	}

}
