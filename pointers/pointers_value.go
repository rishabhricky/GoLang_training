// & gives you the address
// * gives you the value stored at that address

// When you have large chunk of address to pass on.. instead of 
// passing the value, pass the address of the location

package main

import (
	"fmt"
)

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)

	b := &a
	fmt.Println(b)
	fmt.Println(*b)

	*b = 43
	fmt.Println(a)

	x := 0
	foo(&x)
	fmt.Println("address:", &x)
	fmt.Println("Value of a, as value at adddress is changed: ", x)
}

func foo(y *int) {
	fmt.Println("y address", y)
	fmt.Println("y value", *y)

	*y =43
	fmt.Println("y address", y)
	fmt.Println("y value", *y)
	
}
