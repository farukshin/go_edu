package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	cnt := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		cnt[input.Text()]++
	}
	for line, n := range cnt {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
