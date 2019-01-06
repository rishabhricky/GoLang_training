// Methods attached to a type is method set
// Non pointer receiver works with both pointers and non pointers
// Pointer receiver works only with pointers

package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5}
	// info(c)
	info(&c)
}
