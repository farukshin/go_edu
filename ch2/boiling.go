package main

import "fmt"

var boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("Температура кипения = %gC или %gF", c, f)
}
