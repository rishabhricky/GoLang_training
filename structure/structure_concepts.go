package main

import (
	"fmt"
)

type person struct {
	first string
	last string
	age int
}

type secretAgent struct {
	person
	first string // Check name collision
	ltk bool
}

func main(){
	p1 := person{
		first: "James",
		last: "Bond",
		age: 32,
	}

	p2 := person{
		first: "Rishabh",
		last: "Rusia",
		age: 27,
	}

	sa1 := secretAgent{
		person: person{
			first: "Clarke",
			last: "Kent",
			age: 35,
		},
		first: "Ricky", //Collision in name in structure
 		ltk: true,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)

	// Reference properly if there is a collision in names
	fmt.Println(sa1.person.first, sa1.first, sa1.last, sa1.age, sa1.ltk)

	// Anonymous structure (as it doesnt have a name and it helps to keep the code clean)

	p4 := struct{
		first string
		last string
		age int
	}{
		first: "Ricky",
		last: "Rusia",
		age: 27,
	}

	fmt.Println(p4)
}