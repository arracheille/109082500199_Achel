package main

import "fmt"

func main() {
    var n int
	fmt.Print("Bilangan: ")
    fmt.Scan(&n)
	fmt.Print("Faktor: ")
    faktor(1, n)
}

func faktor(i, n int) {
    if i < n {
        if n % i == 0 {
            fmt.Print(i, ", ")
        }
        faktor(i+1, n)
    } else if i == n {
		fmt.Print(i)
	}
}