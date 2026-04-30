package main

import "fmt"

type arrBalita [100]float64

func banyakBalita(T *arrBalita, n *int) {
	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&T[i])
	}
}

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	var j int = 1
	for j < n {
		if arrBerat[j] < *bMin {
			*bMin = arrBerat[j]
		}
		if arrBerat[j] > *bMax {
			*bMax = arrBerat[j]
		}
		j = j + 1
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64 = 0
	for i := 0; i < n; i++ {
		total = total + arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var T arrBalita
	var n int
	var bMin, bMax float64

	banyakBalita(&T, &n)
	hitungMinMax(T, n, &bMin, &bMax)

	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg", rerata(T, n))
}