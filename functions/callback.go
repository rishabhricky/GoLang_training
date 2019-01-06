// Callback is passing the function as an argument
// Functional programming is not recommended in GO

package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := sum(ii...)

	fmt.Println("Sum of all the numbers", s)

	// Now introduce callback concept
	s2 := evenSum(sum, ii...)
	s3 := oddSum(sum, ii...)

	fmt.Println("Sum of even numbers", s2)
	fmt.Println("Sum of odd numbers", s3)

}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func evenSum(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func oddSum(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
