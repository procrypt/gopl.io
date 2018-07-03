package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

func main() {
	count := make(map[string]int)
	for _, filename := range os.Args[1:] {
		bs, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
		}
		count[filename] = 2
		for _, line := range strings.Split(string(bs), "\n") {
			count[line]++
		}
	}
	for line, n :=  range count {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line)
		}
	}

}
