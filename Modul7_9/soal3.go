package main

import "fmt"

const hasil_max int = 100

type hasil_skor [hasil_max]string

func main() {
	var klub_a, klub_b string
	var skor_a, skor_b int
	var hasil hasil_skor
	var n int

	fmt.Print("Klub A : ")
	fmt.Scan(&klub_a)
	fmt.Print("Klub B : ")
	fmt.Scan(&klub_b)

	pertandingan := 1
	for {
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skor_a, &skor_b)

		if skor_a < 0 || skor_b < 0 {
			break
		}

		if skor_a > skor_b {
			hasil[n] = klub_a
		} else if skor_b > skor_a {
			hasil[n] = klub_b
		} else {
			hasil[n] = "Draw"
		}
		n++
		pertandingan++
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Hasil %d : %s\n", i+1, hasil[i])
	}
	fmt.Println("Pertandingan selesai")
}
