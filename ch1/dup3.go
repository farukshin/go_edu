package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	cnt := make(map[string]int)

	for _, file := range os.Args[1:] {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			cnt[line]++
		}
	}

	for line, n := range cnt {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
