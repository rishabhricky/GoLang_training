package main

import (
	"fmt"
)

func main() {
	// Array in go
	x := [5]int{42, 43, 44, 45, 46}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
	fmt.Printf("\n")

	// Slice in go
	y := []int{42, 43, 44, 45, 46, 1, 2, 3, 4}
	for i, v := range y {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", y)
	fmt.Printf("\n")

	// Slice the Slice
	z := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(z[:5])
	fmt.Println(z[5:])
	fmt.Println(z[2:7])
	fmt.Println(z[1:6])
	fmt.Println(z)
	fmt.Println("\n")

	// Slice Manipulation
	w := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	w = append(w, 52)
	fmt.Println(w)
	w = append(w, 53, 54, 55)
	fmt.Println(w)

	q := []int{56, 57, 58, 59, 60}
	w = append(w, q...)
	fmt.Println(w)

	// Delete from a slice
	fmt.Println("\n")
	e := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(e)

	e = append(e[:3], e[6:]...)
	fmt.Println(e)

	// Length and capacity of slice
	fmt.Println("\n")
	//Make creates a slice of given capacity with array so it saves the processing time when size of
	// slice is increased
	states := make([]string, 50, 50)
	states = []string{` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`,
		` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, 
		` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, 
		` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, 
		` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, 
		` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, 
		` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`}

	fmt.Println(states)
	fmt.Println(len(states))
	fmt.Println(cap(states))

	for i:=0; i<len(states); i++ {
		fmt.Println(i, states[i])
	}

	// 2-D string slice

	xs1 := []string{"James", "Bond", "Shaken, not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	fmt.Println(xs1)
	fmt.Println(xs2)

	xxs := [][]string{xs1, xs2}

	for i, xs:= range xxs {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}
