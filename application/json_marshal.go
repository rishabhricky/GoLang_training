//The case of an identifier - lowercase or uppercase, determines whether or not the data can be exported.
package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "Rishabh",
		Last:  "Rusia",
		Age:   27,
	}

	p2 := person{
		First: "Prateek",
		Last:  "Rusia",
		Age:   30,
	}

	people := []person{p1, p2}

	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err	!= nil {
		fmt.Println(err)
	}
	fmt.Printf("%T\n", bs)
	fmt.Printf("%T\n", err)
	fmt.Println(err)
	fmt.Println(bs)
	fmt.Println(string(bs))
}
