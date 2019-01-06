// []uint8 is an alias for slice of bytes []byte
// Understand the concept of tags for getting GO data struct from json
// https://mholt.github.io/json-to-go/

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"Rishabh","Last":"Rusia","Age":27},{"First":"Prateek","Last":"Rusia","Age":30}]`
	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	var people []person

	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	for i, v := range people {
		fmt.Println("\n Person number", i)
		fmt.Println(v.First, v.Last, v.Age)
	}

}
