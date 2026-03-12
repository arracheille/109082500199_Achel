# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[Aqilla Rachel Rabbani] - [109082500199]</p>

## Unguided 

### 1. [Soal]
#### soal1.go

```go
package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul2/output/output-soal2A.png)
##### Penjelasan
Program ini adalah sebuah program dengan konsep pertukaran nilai:
<ol>
  <li>Program memiliki variabel satu, dua tiga dan temp dengan tipe data string.</li>
  <li>Program memberikan input untuk variabel satu, dua dan tiga.</li>
  <li>Setelah memasukkan input, inputan user dibaca oleh program menggunakan Scanln.</li>
  <li>Program menampilkan output awal dari input yang diberi user menggunakan Println.</li>
  <li>Nilai variabel satu disimpan ke variabel temp.</li>
  <li>Nilai variabel dua dipindahkan ke variabel satu.</li>
  <li>Nilai variabel tiga dipindahkan ke variabel dua.</li>
  <li>Nilai yang tadi disimpan di temp (yaitu nilai asli satu) dipindahkan ke variabel tiga. </li>
  <li>Program menampilkan output akhir dari variabel yang sudah dipindah menggunakan Println.</li>
</ol>
