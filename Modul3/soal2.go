package main

import "fmt"

func f (x int) int {
	fx := x * x
	return fx
}

func g (x int) int {
	gx := x - 2
	return gx
}

func h (x int) int {
	hx := x + 1
	return hx
}

func main() {
	var a, b, c int

	fmt.Print("Masukkan nilai variabel a, b, c: ")
	fmt.Scan(&a, &b, &c)

	fmt.Println(f(g(h(a))))
	fmt.Println(g(h(f(b))))
	fmt.Print(h(f(g(c))))
}