package main

import "fmt"

func main() {
	var n int
	fmt.Print("Baris: ")
	fmt.Scan(&n)
	baris(1, n)
}

func baris(i, n int) {
	if i <= n {
		bintang(i)
		fmt.Println()
		baris(i+1, n)
	}
}

func bintang(jumlah int) {
	if jumlah > 0 {
		fmt.Print("*")
		bintang(jumlah - 1)
	}
}