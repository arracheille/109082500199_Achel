package main

import "fmt"

func main(){
	var n int

	fmt.Print("Masukkan nilai variabel n: ")
	fmt.Scan(&n)

	fmt.Print("n: ")
	index(0, n)
	fmt.Print("\nSn: ")
	deret(0, n)

	// pada soal tertulis deretnya sampai 10, tapi biar safe saya pakai variabel n, dibawah ini kodenya kalo sampe 10 doang
	// fmt.Print("n: ")
	// index(0, 10)
	// fmt.Print("\nSn: ")
	// deret(0, 10)
}

func index(x, n int){
	if x == n {
		fmt.Print(x)
	} else {
		fmt.Print(x, " ")
		index(x+1, n)
	}
}

func deret(x, n int){
	if x == n {
		fmt.Print(fibonacci(x))
	} else {
		fmt.Print(fibonacci(x), " ")
		deret(x+1, n)
	}
}

func fibonacci(n int) int {
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    } else {
        return fibonacci(n-1) + fibonacci(n-2)
    }
}