package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, "-- sceretAgent speaks")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, "-- person speaks")
}

type human interface {
	speak()
}

// interface helps to create a common way to handle multiple types

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I was passed into bar and type person", h.(person).first)
	case secretAgent:
		fmt.Println("I was passed into bar and type secretAgent", h.(secretAgent).first)

	fmt.Println("I was passed into bar", h)
	}
}

func main() {
	sa1 := secretAgent{
		person: person{
			"Rishabh",
			"Rusia",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Prateek",
			"Rusia",
		},
		ltk: false,
	}

	p1 := person{
		first: "Dr.",
		last:  "Yes",
	}

	fmt.Println(sa1)
	fmt.Println(sa2)
	fmt.Println(p1)

	sa1.speak()
	sa2.speak()
	p1.speak()

	bar(sa1)
	bar(sa2)
	bar(p1)
}
