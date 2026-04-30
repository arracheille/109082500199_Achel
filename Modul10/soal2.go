package main

import "fmt"

type arrIkan [1000]float64
type arrWadah [1000]float64

func inputIkan(ikan *arrIkan, x *int, y *int) {
	fmt.Print("Masukkan jumlah ikan (x) dan isi per wadah (y): ")
	fmt.Scan(x, y)
	for i := 0; i < *x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}
}

func hitungWadah(ikan arrIkan, x int, y int, wadah *arrWadah) int {
	jumlahWadah := (x + y - 1) / y
	
	for w := 0; w < jumlahWadah; w++ {
		var total float64 = 0
		for i := w * y; i < (w+1)*y && i < x; i++ {
			total = total + ikan[i]
		}
		wadah[w] = total
	}
	return jumlahWadah
}

func rataWadah(wadah arrWadah, jumlahWadah int) float64 {
	var total float64 = 0
	for i := 0; i < jumlahWadah; i++ {
		total = total + wadah[i]
	}
	return total / float64(jumlahWadah)
}

func beratPerWadah(jumlahWadah *int, wadah *arrWadah){
	fmt.Print("Total berat per wadah: ")
	for i := 0; i < *jumlahWadah; i++ {
		fmt.Printf("%.2f ", wadah[i])
	}
}

func main() {
	var ikan arrIkan
	var wadah arrWadah
	var x, y int

	inputIkan(&ikan, &x, &y)

	jumlahWadah := hitungWadah(ikan, x, y, &wadah)

	beratPerWadah(&jumlahWadah, &wadah)

	fmt.Printf("\nRata-rata berat per wadah: %.2f", rataWadah(wadah, jumlahWadah))
}