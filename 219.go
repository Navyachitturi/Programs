package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martins`, `women`},
		"moneypenny_miss": {`James Bond`, `Literature`, `Computer science`},
		"no_dr":           {`being evil`, `Ice cream`, `Sunsets`},
	}
	// fmt.Println(m)

	m["fleming_ian"] = []string{`steaks`, `cigars`, `espionage`}

	for k, v := range m {
		fmt.Println("this is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
