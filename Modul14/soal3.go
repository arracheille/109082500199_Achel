package main

import "fmt"

const max = 1000001

type bilangan [max]int

func urutkan(angka *bilangan, n int) {
	var i, j, temp int
	for i = 1; i <= n-1; i+=1 {
		j = i
		temp = angka[j]
		for j > 0 && temp < angka[j-1] {
			angka[j] = angka[j-1]
			j = j - 1
		}
		angka[j] = temp
	}
}

func main() {
	var data bilangan
	var n, nilai int

	n = 0
	fmt.Println("Masukkan data angka bilangan bulat: ")
	fmt.Scan(&nilai)

	for nilai != -5313 {
		if nilai == 0 {
			urutkan(&data, n)
			if n%2 == 1 {
				fmt.Printf("Median dari %d data adalah: %d\n", n, data[n/2])
			} else {
				fmt.Printf("Median dari %d data adalah: %d\n", n, (data[n/2-1]+data[n/2])/2)
			}
		} else {
			data[n] = nilai
			n++
		}
		fmt.Scan(&nilai)
	}
}