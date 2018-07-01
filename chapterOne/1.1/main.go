package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[:])
	//fmt.Println(strings.Join(os.Args[1:], " "))
}
