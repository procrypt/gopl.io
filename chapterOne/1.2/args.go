package main

import (
	"os"
	"fmt"
)

func main() {
	s, sep := " ", " "
	for key, value := range os.Args[1:] {
		s += sep + value
		sep = " "
		fmt.Println(key, value)
	}
	fmt.Println(s)
}
