package main

import "fmt"

func main() {
	var freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC\n", freezingF, ftoC(freezingF))
	fmt.Printf("%gF = %gC\n", boilingF, ftoC(boilingF))

}
func ftoC(f float64) float64 {
	return (f - 32) * 5 / 9
}
