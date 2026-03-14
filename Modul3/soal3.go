package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var (
		cx1, cy1, r1 float64
		cx2, cy2, r2 float64
		x, y float64
		output string
	)

	fmt.Print("Masukkan koordinat titik pusat x, y dan radius dari lingkaran 1: ")
	fmt.Scanln(&cx1, &cy1, &r1)
	fmt.Print("Masukkan koordinat titik pusat x, y dan radius dari lingkaran 2: ")
	fmt.Scanln(&cx2, &cy2, &r2)
	fmt.Print("Masukkan koordinat titik sembarang (x, y): ")
	fmt.Scan(&x, &y)

	titik1 := didalam(cx1, cy1, r1, x, y)
	titik2 := didalam(cx2, cy2, r2, x, y)

	if titik1 && titik2 {
		output = "Titik di dalam lingkaran 1 dan 2"
	} else if titik1 {
		output = "Titik di dalam lingkaran 1"
	} else if titik2 {
		output = "Titik di dalam lingkaran 2"
	} else {
		output = "Titik di luar lingkaran 1 dan 2"
	}

	fmt.Print(output)
}