package main

import (
	"fmt"
)

const (
	a = 2017 + iota
	b = 2017 + iota
	c = 2017 + iota
	d = 2017 + iota
)

const (
	i = iota
	j 
	k 
)	


func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
}