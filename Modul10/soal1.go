package main

import "fmt"

type anak_kelinci [1000]float64

func jumlah(jumlah_kelinci *anak_kelinci, n *int) {
	fmt.Print("Masukkan jumlah kelinci: ")
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Printf("Berat kelinci ke-%d (kg): ", i+1)
		fmt.Scan(&jumlah_kelinci[i])
	}
}

func berat(jumlah_kelinci anak_kelinci, n int) int {
	var i int = 0
	var j int = 1
	for j < n {
		if jumlah_kelinci[i] < jumlah_kelinci[j] {
			i = j
		}
		j = j + 1
	}
	return i
}

func ringan(jumlah_kelinci anak_kelinci, n int) int {
    var i int = 0
    var j int = 1
    for j < n { 
        if jumlah_kelinci[i] > jumlah_kelinci[j] {
            i = j
        }
        j = j + 1
    }
    return i
}

func main() {
	var jumlah_kelinci anak_kelinci
	var n int

	jumlah(&jumlah_kelinci, &n)

	terberat := berat(jumlah_kelinci, n)
	terringan := ringan(jumlah_kelinci, n)

	fmt.Println("Kelinci paling berat adalah kelinci dengan berat:", jumlah_kelinci[terberat])
	fmt.Println("Kelinci paling ringan adalah kelinci dengan berat:", jumlah_kelinci[terringan])
}