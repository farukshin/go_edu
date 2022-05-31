package main

import (
	"bufio"
	"fmt"
	"os"
)

func cntLine(f *os.File, cnt map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		cnt[input.Text()]++
	}
}

func main() {
	cnt := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		cntLine(os.Stdin, cnt)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			cntLine(f, cnt)
			f.Close()
		}
	}

	for line, n := range cnt {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
