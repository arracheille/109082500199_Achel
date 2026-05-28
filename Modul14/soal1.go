package main

import "fmt"

const banyakRumah = 1000000

type kerabat [banyakRumah]int

func urutNomor(nomor *kerabat, n int) {
	var t, i, j, idx_min int
	for i = 1; i <= n-1; i+=1 {
		idx_min = i - 1
		j = i
		for j < n {
			if nomor[idx_min] > nomor[j] {
				idx_min = j
			}
			j = j + 1
		}
		t = nomor[idx_min]
		nomor[idx_min] = nomor[i-1]
		nomor[i-1] = t
	}
}

func main() {
	var n, m int
	var nomorRumah kerabat
	var i, j, k int

	fmt.Print("Masukkan banyaknya daerah kerabat Hercules: ")
	fmt.Scan(&n)
	
	for i = 0; i < n; i++ {
		fmt.Printf("\nMasukkan nomor rumah di daerah ke-%d: ", i+1)
		fmt.Scan(&m)
		for j = 0; j < m; j++ {
			fmt.Scan(&nomorRumah[j])
		}
		
		urutNomor(&nomorRumah, m)
		fmt.Printf("Nomor rumah terurut (%d rumah): ", m)

		for k = 0; k < m; k++ {
			if k > 0 {
				fmt.Print(" ")
			}
			fmt.Print(nomorRumah[k])
		}
		fmt.Println()
	}
}