package main

import "fmt"

func main() {
	var parsel, kg, gr, harga_kg, harga_gr, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&parsel)

	kg = parsel / 1000
	gr = parsel % 1000

	harga_kg = kg * 10000

	if gr >= 500 {
		harga_gr = gr * 5
	} else if gr > 0 {
		harga_gr = gr * 15
	} else {
		harga_gr = 0
	}

	fmt.Printf("Detail berat: %d kg + %d gr", kg, gr)
	fmt.Printf("\nDetail biaya: Rp. %d + Rp. %d", harga_kg, harga_gr)

	if parsel > 10000 {
		total = harga_kg
	} else {
		total = harga_kg + harga_gr
	}

	fmt.Printf("\nTotal biaya: Rp. %d", total)
}