// Every time a new variable is assigned to the calling function, it
// gets created in a new location and new variales are created for it
package main

import (
	"fmt"
)

func main() {

	{
		var y int
		fmt.Println(y)
	}

	// We can create a code block inside code block to limit thye access of 
	// variable


	a := incrementor()
	b := incrementor()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

}

func incrementor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
