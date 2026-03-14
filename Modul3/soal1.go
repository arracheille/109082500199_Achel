package main

import "fmt"

func factorial (n int) int {
	j := 1
	for i := 1; i <= n; i++{
		j *= i
	}
	return j
}

func permutation (n,r int) int {
	p := factorial(n) / factorial(n-r)
	return p
}

func combination (n,r int) int {
	c := factorial(n) / (factorial(r) * factorial(n-r))
	return c
}

func main() {
	var a, b, c, d int

	fmt.Print("Masukkan nilai variabel a, b, c, d: ")
	fmt.Scan(&a, &b, &c, &d)

	fmt.Print(permutation(a, c), " ")
	fmt.Println(combination(a, c))
	
	fmt.Print(permutation(b, d), " ")
	fmt.Print(combination(b, d))
}