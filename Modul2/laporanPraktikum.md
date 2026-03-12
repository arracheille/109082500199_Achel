# <h1 align="center">Laporan Praktikum Modul 2 - ... </h1>
<p align="center">[Aqilla Rachel Rabbani] - [109082500199]</p>

## Unguided 

### 1. Soal Latihan Modul 2A
#### soal2A.go

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
<ul>
  <li>Program memiliki variabel satu, dua tiga dan temp dengan tipe data string.</li>
  <li>Program memberikan input untuk variabel satu, dua dan tiga.</li>
  <li>Setelah memasukkan input, inputan user dibaca oleh program menggunakan Scanln.</li>
  <li>Program menampilkan output awal dari input yang diberi user menggunakan Println.</li>
  <li>Nilai variabel satu disimpan ke variabel temp.</li>
  <li>Nilai variabel dua dipindahkan ke variabel satu.</li>
  <li>Nilai variabel tiga dipindahkan ke variabel dua.</li>
  <li>Nilai yang tadi disimpan di temp (yaitu nilai asli satu) dipindahkan ke variabel tiga. </li>
  <li>Program menampilkan output akhir dari variabel yang sudah dipindah menggunakan Println.</li>
</ul>

### 2. Soal Latihan Modul 2B
#### soal2B.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_2](https://github.com/arracheille/109082500199_Achel/blob/main/Modul2/output/output-soal2B.png)
##### Penjelasan
Program ini adalah sebuah program dengan konsep pertukaran nilai:
<ul>
  <li>Program memiliki variabel gelas1, gelas2, gelas3 dan gelas4 dengan tipe data string.</li>
  <li>Program membuat variabel baru bernama berhasil dan diberi nilai true (boolean).</li>
  <li>Program membuat perulangan for dengan:
	<ol type="a">
		<li>inisiasi, yaitu program membuat variabel baru bernama percobaan menggunakan “:=” dan diberi nilai 1.</li>
		<li>kondisinya adalah nilai variabel percobaan kurang dari sama dengan 5.</li>
		<li>updatenya adalah percobaan++ (Post-increment).</li>
	</ol>
  </li>
  <li>Lalu kode yang dibuat di dalam for adalah:
	<ol type="a">
		<li>Text perintah untuk memasukkan nilai variabel gelas1, gelas2, gelas3 dan gelas4 menggunakan fmt.Printf (print format), yang bertuliskan “Percobaan %d: ”. Saat di-outputkan, %d nilainya akan berubah menjadi nilai variabel percobaan. </li>
		<li>Input untuk nilai variabel gelas1, gelas2, gelas3 dan gelas4. Setelah memasukkan input, inputan user dibaca oleh program menggunakan fmt.Scan dengan &gelas1, &gelas2, &gelas3 dan &gelas4, lalu nilai inputan dimasukkan ke dalam masing-masing variabel secara berurutan.</li>
		<li>Kondisi if, yaitu jika nilai variabel gelas1 tidak sama dengan “merah” atau nilai variabel gelas2 tidak sama dengan “kuning” atau nilai variabel gelas3 tidak sama dengan “hijau” atau nilai variabel gelas4 tidak sama dengan “ungu”, maka nilai variabel berhasil diubah menjadi false.</li>
	</ol>
  </li>
  <li>Terakhir, ada output menggunakan fmt.Print yang bertuliskan “BERHASIL: (nilai variabel berhasil)”, nilai variabel berhasil bisa bertulis true atau false.</li>
</ul>

### 3. Soal Latihan Modul 2C
#### soal2C.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_3](https://github.com/arracheille/109082500199_Achel/blob/main/Modul2/output/output-soal2C.png)
##### Penjelasan
Program ini adalah sebuah program dengan konsep pertukaran nilai:
<ul>
  <li>Program memiliki variabel parsel, kg, gr, harga_kg, harga_gr dan total dengan tipe data int (integer).</li>
  <li>Program menampilkan text perintah untuk memasukkan nilai variabel parsel kepada user menggunakan fmt.Print.</li>
  <li>Program memberikan input untuk nilai variabel parsel.</li>
  <li>Setelah memasukkan input, inputan user dibaca oleh program menggunakan fmt.Scan, lalu mencari input source variabel orang menggunakan &parsel.</li>
  <li>Program memberikan rumus untuk variabel kg, yaitu nilai variabel parsel dibagi 1000. Variabel kg adalah berat parsel dalam kg.</li>
  <li>Program memberikan rumus untuk variabel gr, yaitu nilai variabel parsel di-modulus 1000. Variabel gr adalah sisa berat parsel dalam gram.</li>
  <li>Program memberikan rumus untuk variabel harga_kg, yaitu nilai kg dikali 10000.</li>
  <li>Program membuat kondisi if-else untuk rumus harga dalam gram dengan:
	<ol type="a">
		<li><strong>kondisi 1 (if)</strong>, jika nilai variabel gr, lebih dari sama dengan 500, maka rumus untuk variabel harga_gr adalah nilai variabel gr dikali 5.</li>
		<li><strong>kondisi 2 (else if)</strong>, jika nilai variabel gr, lebih dari 0, maka rumus untuk variabel harga_gr adalah nilai variabel gr dikali 15.</li>
		<li><strong>kondisi 3 (else)</strong>, jika nilai variabel gr tidak memenuhi kedua kondisi sebelumnya, maka nilai variabel harga_gr adalah 0.</li>
	</ol>
  </li>
  <li>Program menampilkan output detail berat berupa kg dan gram (nilai dari variabel kg dan gr).</li>
  <li>Program menampilkan output detail biaya berupa harga parcel yang kg dan yang gram (nilai dari variabel harga_kg dan harga_gr).</li>
  <li>Program membuat kondisi if-else untuk total harga dengan:
	<ol type="a">
		<li><strong>kondisi 1 (if)</strong>, jika nilai variabel parsel lebih dari 10000, maka nilai variabel total adalah nilai variabel harga_kg.</li>
		<li><strong>kondisi 2 (else)</strong>, jika variabel parcel tidak memenuhi kondisi pertama, maka akan menggunakan kondisi kedua, yaitu  nilai variabel total adalah variabel harga_kg ditambah harga_gr.</li>
	</ol>
  </li>
  <li>Program menampilkan output terakhir yaitu total biaya (nilai dari variabel total).</li>
  <li>Semua output variabel menggunakan fmt.Printf:
	<ol type="a">
		<li>fmt.Printf yaitu print format.</li>
		<li>menampilkan text menggunakan (“”).</li>
		<li>%d (format bilangan bulat) untuk menampilkan value setiap variabel.</li>
		<li>\n untuk line baru</li>
	</ol>
  </li>
</ul>


