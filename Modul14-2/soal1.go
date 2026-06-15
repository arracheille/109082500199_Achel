package main

import "fmt"

type arrInt [1000001]int

func insertionSort1(T *arrInt, n int) {
	var temp, i, j int

	for i = 1; i <= n-1; i+=1 {
		j = i
		temp = T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
	}
}

func main() {
	var T arrInt
	var n, x int
	var selesai bool

	n = 0
	selesai = false

	fmt.Print("Masukkan data nomor: ")
    for !selesai {
        fmt.Scan(&x)
        
        if x < 0 {
            selesai = true
        } else {
            T[n] = x
            n = n + 1
        }
    }

	insertionSort1(&T, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(T[i])
	}
	fmt.Println()

	if n <= 1 {
		fmt.Println("Data berjarak tidak tetap")
	} else {
		jarak := T[1] - T[0]
		tetap := true
		for j := 2; j < n; j+=1 {
			if T[j]-T[j-1] != jarak {
				tetap = false
			}
		}
		if tetap {
			fmt.Println("Data berjarak", jarak)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}