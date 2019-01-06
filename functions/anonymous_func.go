package main

import (
	"fmt"
)

func main() {
	foo()

	func() {
		fmt.Println("inside anonymous func")
	}()

	func(x int) {
		fmt.Println("Meaning of Life: ", x)
	}(42)

	// functions are first class citizens and can be used as types
	f := func() {
		fmt.Println("Hello from func expressions")
	}
	f()

	g := func(x int) {
		fmt.Println("The year when I was born: ", x)
	}
	g(1991)

	// Returning a func

	x := bar()
	// Checking the type of return type
	fmt.Printf("%T\n", x)

	i := x()
	fmt.Println(i)
	fmt.Println(bar()())

}

func foo() {
	fmt.Println("Inside foo")
}

func bar() func() int {
	return func() int {
		return 451
	}
}
