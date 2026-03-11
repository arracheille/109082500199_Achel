package main

import "fmt"

func ubahNilai(x int) {
	x = x * 2
}

func main() {
	var angka int = 5
	ubahNilai(angka)
	fmt.Print(angka)
}