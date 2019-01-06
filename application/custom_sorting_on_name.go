package main

import (
	"fmt"
	"sort"
)

// If the type Interface interface is implemented, these three methods shoud be there, sort uses them and helps to sort
// These three methods implements this interface from sort
type Person struct {
	Name string
	Age  int
}

type byName []Person

func (bn byName) Len() int           { return len(bn) }
func (bn byName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn byName) Less(i, j int) bool { return bn[i].Name < bn[j].Name }

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Moneypenny", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 56}

	people := []Person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(byName(people))
	fmt.Println(people)
}
