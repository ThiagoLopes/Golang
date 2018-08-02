package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// with range
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)

	// Best way
	fmt.Println(strings.Join(os.Args[1:], " "))
}
