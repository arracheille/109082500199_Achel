package main

import "fmt"

func hitungSkor(soal *int, skor *int) {
	var menit int
	*soal = 0
	*skor = 0
	for i := 1; i <= 8; i++ {
		fmt.Scan(&menit)
		if menit <= 300 {
			*soal++
			*skor += menit
		}
	}
}

func main() {
	var nama, pemenang string
	var soal, skor, soal_max, skor_menang int

	skor_menang = 999999

	fmt.Println("Masukkan nama, dan 8 nilai menit setiap soal:")
	fmt.Scan(&nama)

	for nama != "Selesai" {
		hitungSkor(&soal, &skor)

		if soal > soal_max || (soal == soal_max && skor < skor_menang) {
			soal_max = soal
			skor_menang = skor
			pemenang = nama
		}

		fmt.Scan(&nama)
	}

	fmt.Println(pemenang, soal_max, skor_menang)
}