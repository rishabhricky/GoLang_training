package main

import (
	"fmt"
)

func main() {
	// Map as key of string and values as slice(list) of string 
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v{
			fmt.Println("\t", i, v2)
		}
	}

	// Add the record to the map
	m[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`}

	for k, v := range m {
		fmt.Println("This is the record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

	// Delete the record from map
	delete(m, `no_dr`)
	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

	// Other way
	r := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(r)

	fmt.Println(r["James"])

	fmt.Println(r["Barnabas"])

	v, ok := r["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Barnabas"]; ok {
		fmt.Println(v)
	}
}