package main

import "fmt"

func factorial (n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++{
		*hasil *= i
	}
}

func permutation (n,r int, hasil *int) {
	var faktorial_n, faktorial_nr int
	factorial(n, &faktorial_n)
	factorial(n-r, &faktorial_nr)
	*hasil = faktorial_n / faktorial_nr
}

func combination (n,r int, hasil *int) {
	var faktorial_n, faktorial_r, faktorial_nr int
	factorial(n, &faktorial_n)
	factorial(r, &faktorial_r)
	factorial(n-r, &faktorial_nr)
	*hasil = faktorial_n / (faktorial_r * faktorial_nr)
}

func main() {
	var a, b, c, d int
	var p1, c1, p2, c2 int

	fmt.Print("Masukkan nilai variabel a, b, c, d: ")
	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &p1)
	combination(a, c, &c1)
	fmt.Println(p1, c1)
	
	permutation(b, d, &p2)
	combination(b, d, &c2)
	fmt.Print(p2, c2)
}