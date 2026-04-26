package main

import (
	"fmt"
	"math"
)

const array_max int = 100

type tabelInt [array_max]int

func array(t *tabelInt, n *int) {
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&t[i])
	}
}

func tampilSemua(t tabelInt, n int) {
	fmt.Print("Semua elemen: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", t[i])
	}
	fmt.Println()
}

func tampilGanjil(t tabelInt, n int) {
	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Printf("%d ", t[i])
	}
	fmt.Println()
}

func tampilGenap(t tabelInt, n int) {
	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Printf("%d ", t[i])
	}
	fmt.Println()
}

func tampilKelipatanX(t tabelInt, n, x int) {
	fmt.Printf("Indeks kelipatan %d: ", x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Printf("%d ", t[i])
		}
	}
	fmt.Println()
}

func hapusElemen(t *tabelInt, n *int, idx int) {
	for i := idx; i < *n-1; i++ {
		t[i] = t[i+1]
	}
	*n--
}

func rataRata(t tabelInt, n int) float64 {
	sum := 0
	for i := 0; i < n; i++ {
		sum += t[i]
	}
	return float64(sum) / float64(n)
}

func stdDeviasi(t tabelInt, n int) float64 {
	mean := rataRata(t, n)
	sum := 0.0
	for i := 0; i < n; i++ {
		diff := float64(t[i]) - mean
		sum += diff * diff
	}
	return math.Sqrt(sum / float64(n))
}

func frekuensi(t tabelInt, n, cari int) int {
	freq := 0
	for i := 0; i < n; i++ {
		if t[i] == cari {
			freq++
		}
	}
	return freq
}

func main() {
	var tab tabelInt
	var n, x, idx, cari int

	array(&tab, &n)

	tampilSemua(tab, n)

	tampilGanjil(tab, n)

	tampilGenap(tab, n)

	fmt.Print("Masukkan nilai x untuk kelipatan: ")
	fmt.Scan(&x)
	tampilKelipatanX(tab, n, x)

	fmt.Print("Masukkan indeks yang akan dihapus: ")
	fmt.Scan(&idx)
	hapusElemen(&tab, &n, idx)
	fmt.Print("Array setelah penghapusan: ")
	tampilSemua(tab, n)

	fmt.Printf("Rata-rata: %.2f\n", rataRata(tab, n))

	fmt.Printf("Standar deviasi: %.2f\n", stdDeviasi(tab, n))

	fmt.Print("Cari frekuensi bilangan: ")
	fmt.Scan(&cari)
	fmt.Printf("Frekuensi %d: %d\n", cari, frekuensi(tab, n, cari))
}