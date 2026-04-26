package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	pusat titik
	r int
}

func jarak(p, q titik) float64 {
	dx := float64(p.x - q.x)
	dy := float64(p.y - q.y)
	return math.Sqrt(dx*dx + dy*dy)
}

func didalam(c lingkaran, p titik) bool {
	return jarak(c.pusat, p) <= float64(c.r)
}

func main() {
	var (
		c1, c2 lingkaran
		p titik
		output string
	)

	fmt.Print("Masukkan koordinat titik pusat x, y dan radius dari lingkaran 1: ")
	fmt.Scanln(&c1.pusat.x, &c1.pusat.y, &c1.r)
	fmt.Print("Masukkan koordinat titik pusat x, y dan radius dari lingkaran 2: ")
	fmt.Scanln(&c2.pusat.x, &c2.pusat.y, &c2.r)
	fmt.Print("Masukkan koordinat titik sembarang (x, y): ")
	fmt.Scan(&p.x, &p.y)

	titik1 := didalam(c1, p)
	titik2 := didalam(c2, p)

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