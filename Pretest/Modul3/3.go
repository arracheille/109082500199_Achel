package main

import "fmt"

func bagi(a, b float64) float64 {
	return a / b
}

func main() {
	x, y := 5, 2
	fmt.Print(bagi(x, y))
}