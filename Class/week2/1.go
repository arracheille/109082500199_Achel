package main

import "fmt"

func main(){
	var (
		satu, dua, tiga string
		temp string
	)

	fmt.Print("Masukkan input string: ")
	fmt.Scan(&satu)
	fmt.Print("Masukkan input string: ")
	fmt.Scan(&dua)
	fmt.Print("Masukkan input string: ")
	fmt.Scan(&tiga)
	fmt.Println("Output awal = " + satu + "" + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + "" + dua + " " + tiga)
}