package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello")
	fmt.Fprintln(os.Stdout, "Hello") // Println implements this
	io.WriteString(os.Stdout, "Hello")
}
