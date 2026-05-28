package main

import "fmt"

const rumah = 1000000

type rumahKerabat [rumah]int

func urutKecil(nomor *rumahKerabat, n int) {
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

func urutBesar(nomor *rumahKerabat, n int) {
	var t, i, j, idx_max int
	for i = 1; i <= n-1; i+=1 {
		idx_max = i - 1
		j = i
		for j < n {
			if nomor[idx_max] < nomor[j] {
				idx_max = j
			}
			j = j + 1
		}
		t = nomor[idx_max]
		nomor[idx_max] = nomor[i-1]
		nomor[i-1] = t
	}
}

func main() {
	var n, m int
	var ganjil, genap rumahKerabat
	var i, j, k, l int
	var nGanjil, nGenap, val int

	fmt.Print("Masukkan banyaknya daerah kerabat Hercules: ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Printf("\nMasukkan nomor rumah di daerah ke-%d: ", i+1)
		fmt.Scan(&m)

		nGanjil = 0
		nGenap = 0

		for j = 0; j < m; j++ {
			fmt.Scan(&val)
			if val%2 != 0 {
				ganjil[nGanjil] = val
				nGanjil++
			} else {
				genap[nGenap] = val
				nGenap++
			}
		}

		urutKecil(&ganjil, nGanjil)
		urutBesar(&genap, nGenap)

		fmt.Printf("Nomor rumah terurut (%d ganjil, %d genap): ", nGanjil, nGenap)

		for k = 0; k < nGanjil; k++ {
			if k > 0 {
				fmt.Print(" ")
			}
			fmt.Print(ganjil[k])
		}

		for l = 0; l < nGenap; l++ {
			if nGanjil > 0 || l > 0 {
				fmt.Print(" ")
			}
			fmt.Print(genap[l])
		}
		fmt.Println()
	}
}