package main

import "fmt"

func main() {
	var gelas1, gelas2, gelas3, gelas4 string
	
	berhasil := true

	for percobaan := 1; percobaan <= 5; percobaan++ {
		fmt.Printf("Percobaan %d: ", percobaan)
		fmt.Scan(&gelas1, &gelas2, &gelas3, &gelas4)
		if gelas1 != "merah" || gelas2 != "kuning" || gelas3 != "hijau" || gelas4 != "ungu" {
			berhasil = false
		}
	}

	fmt.Print("BERHASIL: ", berhasil)
}