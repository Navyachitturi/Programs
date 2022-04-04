package main

import "fmt"

func main() {
	x := make([]string, 50, 50)
	fmt.Println("third time with x")
	fmt.Println(len(x))
	fmt.Println(cap(x))
	states := []string{`Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`, `Connecticut`, `Delaware`, `Florida`, `Georgia`, `Hawaii`, `Idaho`, `lllinois`, `indiana`, `lowa`, `Kansas`, `Kentucky`, `Louisiana`, `Maine`, `Maryland`, `Massachusetts`, `Michigan`, `Minnesota`, `Mississippi`, `Missouri`, `Montana`, `Nebraska`, `Nevada`, `New Hampshire`, `New Jersey`,
		`New Mexico`, `New York`, `North carolina`, `North Dakota`, `Ohio`, `Oklahoma`, `Oregon`, `Pennsylvania`, `Rhode Island`, `south carolina`, `South Dakota`, `Tennessee`, `Texas`, `Utah`, `Vermont`,
		`Virginia`, `Washington`, `West virginia`, `wisconsin`, `Wyoming`}

	for i, v := range states {
		x[i] = v
	}

	fmt.Println("fourth time with x")
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

}
