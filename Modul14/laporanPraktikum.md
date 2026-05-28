# <h1 align="center">Laporan Praktikum Modul 14 - Soal Latihan Modul 14 </h1>

<p align="center">Aqilla Rachel Rabbani - 109082500199</p>

## Unguided

### 1. Hercules, preman terkenal seantero ibukota, memiliki kerabat di banyak daerah. Tentunya Hercules sangat suka mengunjungi semua kerabatnya itu.

### Diberikan masukan nomor rumah dari semua kerabatnya di suatu daerah, buatlah program rumahkerabat yang akan menyusun nomor-nomor rumah kerabatnya secara terurut membesar menggunakan algoritma selection sort. 

### Masukan dimulai dengan sebuah integer n (0 < n < 1000), banyaknya daerah kerabat Hercules tinggal. Isi n baris berikutnya selalu dimulai dengan sebuah integer m (0 < m < 1000000) yang menyatakan banyaknya rumah kerabat di daerah tersebut, diikuti dengan rangkaian bilangan bulat positif, nomor rumah para kerabat.

### Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar di masing- masing daerah.

#### soal1.go

```go
package main

import "fmt"

const banyakRumah = 1000000

type kerabat [banyakRumah]int

func urutNomor(nomor *kerabat, n int) {
	var t, i, j, idx_min int
	for i = 1; i <= n-1; i+=1 {
		idx_min = i - 1
		j = i
		for j < n {
			if nomor[idx_min] > nomor[j] {
				idx_min = j
			}
			j = j + 1
		}
		t = nomor[idx_min]
		nomor[idx_min] = nomor[i-1]
		nomor[i-1] = t
	}
}

func main() {
	var n, m int
	var nomorRumah kerabat
	var i, j, k int

	fmt.Print("Masukkan banyaknya daerah kerabat Hercules: ")
	fmt.Scan(&n)
	
	for i = 0; i < n; i++ {
		fmt.Printf("\nMasukkan nomor rumah di daerah ke-%d: ", i+1)
		fmt.Scan(&m)
		for j = 0; j < m; j++ {
			fmt.Scan(&nomorRumah[j])
		}
		
		urutNomor(&nomorRumah, m)
		fmt.Printf("Nomor rumah terurut (%d rumah): ", m)

		for k = 0; k < m; k++ {
			if k > 0 {
				fmt.Print(" ")
			}
			fmt.Print(nomorRumah[k])
		}
		fmt.Println()
	}
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul14/output/output-soal1.png)

##### Penjelasan
Program ini mengurutkan nomor rumah kerabat Hercules di setiap daerah secara ascending menggunakan algoritma Selection Sort. Berikut penjelasan yang lebih lanjut:

<ul>
    <li>Tipe array <strong>kerabat</strong>
    <br/>Program memiliki array bernama kerabat dengan kapasitas banyakRumah sebesar 1000000 dan tipe data integer.
    </li>
    <li><strong>Di dalam func main()</strong>
        <ol>
            <li>Program memiliki variabel n dan m dengan tipe data integer</li>
            <li>Program memiliki variabel nomorRumah dengan tipe data kerabat, yang berarti nomorRumah adalah array.</li>
            <li>Program memiliki variabel i, j, dan k dengan tipe data integer</li>
            <li>User diberi perintah untuk menginputkan variabel n menggunakan fmt.Scan, perintah tersebut ditulis menggunakan fmt.Print.</li>
            <li>Program membuat perulangan for dengan:
                <ol type="a">
                    <li><strong>inisiasi</strong>, yaitu variabel i diberi nilai 0.</li>
                    <li><strong>kondisinya</strong> adalah nilai variabel i kurang dari variabel n.</li>
                    <li><strong>updatenya</strong> adalah i++ (Post-increment).</li>
                </ol>
            </li>
            <li>Kode di dalam perulangan for:
                <ol type="a">
                    <li>Fmt.Scan untuk memasukkan banyaknya nilai variabel m, dengan angka awal adalah banyaknya data variabel m.</li>
                    <li>
                        <li>Perulangan for:
                            <ol>
                                <li><strong>inisiasi</strong>, yaitu variabel j diberi nilai 0.</li>
                                <li><strong>kondisinya</strong> adalah nilai variabel j kurang dari variabel n.</li>
                                <li><strong>updatenya</strong> adalah j++ (Post-increment).</li>
                                <li><strong>isinya</strong> adalah fmt.Scan dengan untuk memasukkan array nomorRumah, nilai dari setiap array urut berdasarkan index variabel j.</li>
                            </ol>
                        </li>
                    </li>
                </ol>
            </li>
        </ol>
    </li>
    <li><strong>Prosedur dan fungsi</strong>
        <ol>
            <li>Prosedur <strong>jumlah</strong>
            </br>Prosedur ini adalah prosedur untuk menampilkan input secara berulang berdasarkan banyaknya jumlah kelinci. Prosedur ini memiliki parameter pass by reference yaitu jumlah_kelinci dengan tipe data anak_kelinci dan n dengan tipe data integer. Kedua parameter diakses menggunakan simbol bintang (*) karena merupakan pointer ke tipe datanya. Di dalam prosedur ini terdapat:
                <ol>
                    <li>Teks menggunakan fmt.Print untuk menampilkan teks perintah memasukkan jumlah kelinci ke user.</li>
                    <li>Input menggunakan fmt.Scan, inputan dimasukkan ke variabel n.</li>
                    <li>Perulangan for dengan inisiasi membuat variabel baru bernama i menggunakan := dan diberi value 0, kondisi variabel i lebih kecil dari value variabel n, update i++ (Post-increment).</li>
                    <li>Isi dari perulangan ini adalah teks yang menampilkan index berat kelinci menggunakan fmt.Printf dan format %d dari variabel i+1. Dan juga input menggunakan fmt.Scan untuk array jumlah_kelinci. Lalu inputan dibaca oleh program menggunakan fmt.Scan dan &jumlah_kelinci[i], lalu nilai inputan dimasukkan ke array jumlah_kelinci berdasarkan index ([i]) secara berurutan.</li>
                </ol>
            </li>
            <li>Fungsi <strong>berat</strong>
            </br>Fungsi ini adalah fungsi untuk mencari elemen terbesar dari array jumlah_kelinci. Prosedur ini memiliki parameter pass by value yaitu jumlah_kelinci dengan tipe data anak_kelinci dan n dengan tipe data integer. Tipe data yang dikembalikan adalah integer. Di dalam fungsi ini terdapat:
                <ol>
                    <li>Variabel i dengan tipe data integer dan diberi nilai 0</li>
                    <li>Variabel j dengan tipe data integer dan diberi nilai 1</li>
                    <li>Perulangan for dengan kondisi j lebih kecil dari n.</li>
                    <li>Isi dari perulangan ini adalah if dengan kondisi array jumlah_kelinci dengan index ke-i <strong>lebih kecil</strong> dari array jumlah_kelinci dengan index ke-j. Jika kondisi tersebut terpenuhi maka value variabel i adalah nilai variabel j. Lalu ada juga value variabel j adalah nilai variabel j ditambah 1.</li>
                </ol>
            Return atau nilai yang dikembalikan dari fungsi ini adalah variabel i
            </li>
            <li>Fungsi <strong>ringan</strong>
            </br>Fungsi ini adalah fungsi untuk mencari elemen terkecil dari array jumlah_kelinci. Prosedur ini memiliki parameter pass by value yaitu jumlah_kelinci dengan tipe data anak_kelinci dan n dengan tipe data integer. Tipe data yang dikembalikan adalah integer. Di dalam fungsi ini terdapat:
                <ol>
                    <li>Variabel i dengan tipe data integer dan diberi nilai 0</li>
                    <li>Variabel j dengan tipe data integer dan diberi nilai 1</li>
                    <li>Perulangan for dengan kondisi j lebih kecil dari n.</li>
                    <li>Isi dari perulangan ini adalah if dengan kondisi array jumlah_kelinci dengan index ke-i <strong>lebih besar</strong> dari array jumlah_kelinci dengan index ke-j. Jika kondisi tersebut terpenuhi maka value variabel i adalah nilai variabel j. Lalu ada juga value variabel j adalah nilai variabel j ditambah 1.</li>
                </ol>
            Return atau nilai yang dikembalikan dari fungsi ini adalah variabel i.
            </li>
        </ol>
    </li>
</ul>

### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual.

### Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual.

### Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.

#### soal2.go

```go
package main

import "fmt"

const rumah = 1000000

type rumahKerabat [rumah]int

func urutKecil(nomor *rumahKerabat, n int) {
	var t, i, j, idx_min int
	for i = 1; i <= n-1; i+=1 {
		idx_min = i - 1
		j = i
		for j < n {
			if nomor[idx_min] > nomor[j] {
				idx_min = j
			}
			j = j + 1
		}
		t = nomor[idx_min]
		nomor[idx_min] = nomor[i-1]
		nomor[i-1] = t
	}
}

func urutBesar(nomor *rumahKerabat, n int) {
	var t, i, j, idx_max int
	for i = 1; i <= n-1; i+=1 {
		idx_max = i - 1
		j = i
		for j < n {
			if nomor[idx_max] < nomor[j] {
				idx_max = j
			}
			j = j + 1
		}
		t = nomor[idx_max]
		nomor[idx_max] = nomor[i-1]
		nomor[i-1] = t
	}
}

func main() {
	var n, m int
	var ganjil, genap rumahKerabat
	var i, j, k, l int
	var nGanjil, nGenap, val int

	fmt.Print("Masukkan banyaknya daerah kerabat Hercules: ")
	fmt.Scan(&n)

	for i = 0; i < n; i++ {
		fmt.Printf("\nMasukkan nomor rumah di daerah ke-%d: ", i+1)
		fmt.Scan(&m)

		nGanjil = 0
		nGenap = 0

		for j = 0; j < m; j++ {
			fmt.Scan(&val)
			if val%2 != 0 {
				ganjil[nGanjil] = val
				nGanjil++
			} else {
				genap[nGenap] = val
				nGenap++
			}
		}

		urutKecil(&ganjil, nGanjil)
		urutBesar(&genap, nGenap)

		fmt.Printf("Nomor rumah terurut (%d ganjil, %d genap): ", nGanjil, nGenap)

		for k = 0; k < nGanjil; k++ {
			if k > 0 {
				fmt.Print(" ")
			}
			fmt.Print(ganjil[k])
		}

		for l = 0; l < nGenap; l++ {
			if nGanjil > 0 || l > 0 {
				fmt.Print(" ")
			}
			fmt.Print(genap[l])
		}
		fmt.Println()
	}
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul14/output/output-soal2.png)

##### Penjelasan

Program ini adalah program untuk mencari rata-rata berat ikan dari setiap wadah. Berikut penjelasan yang lebih lanjut:

<ul>
    <li>Tipe array <strong>arrIkan</strong> dan <strong>arrWadah</strong>
        <br/>Program memiliki dua tipe array, arrIkan dan arrWadah, keduanya berkapasitas 1000 dan tipe datanya adalah float64.
    </li>
    <li><strong>Di dalam func main()</strong>
        <ol>
            <li>Program memiliki variabel ikan dengan tipe data arrIkan, yang berarti ikan adalah array untuk menyimpan berat ikan.</li>
            <li>Program memiliki variabel wadah dengan tipe data arrWadah, yang berarti wadah adalah array untuk menyimpan total berat per wadah.</li>
            <li>Program memiliki variabel x dan y dengan tipe data integer.</li>
            <li>Program memanggil prosedur inputIkan dengan parameter aktual &ikan, &x, dan &y.</li>
            <li>Program membuat variabel baru bernama jumlahWadah menggunakan := dengan value fungsi hitungWadah dan argumen fungsi adalah ikan, x, y, dan parameter aktual &wadah.</li>
            <li>Program memanggil prosedur beratPerWadah dengan parameter aktual &jumlahWadah dan &wadah.</li>
            <li>Program memberi output menggunakan fmt.Printf (print format) untuk menampilkan rata-rata berat per wadah dari fungsi rataWadah dengan argumen wadah dan jumlahWadah, dengan format %.2f (2 angka di belakang koma) dan juga "\n" untuk mencetak teksnya di baris baru.</li>
        </ol>
    </li>
    <li><strong>Prosedur dan fungsi</strong>
        <ol>
            <li>Prosedur <strong>inputIkan</strong>
                <br/>Prosedur ini adalah prosedur untuk menerima input data berat ikan. Prosedur ini memiliki parameter pass by reference yaitu ikan dengan tipe data arrIkan, dan x dan y dengan tipe data integer. Ketiga parameter diakses menggunakan simbol bintang (*) karena merupakan pointer ke tipe datanya. Di dalam prosedur ini terdapat:
                <ol>
                    <li>Teks menggunakan fmt.Print untuk menampilkan perintah memasukkan jumlah ikan isi per wadah ke user.</li>
                    <li>Input menggunakan fmt.Scan untuk variabel x dan y, lalu hasil input dimasukkan ke variabel x dan y secara berurutan.</li>
                    <li>Perulangan for dengan inisiasi membuat variabel baru bernama i menggunakan := dan diberi value 0, kondisi variabel i lebih kecil dari value variabel x, update i++ (Post-increment).</li>
                    <li>Isi dari perulangan ini adalah teks yang menampilkan index berat ikan menggunakan fmt.Printf dan format %d dari variabel i+1. Dan juga input menggunakan fmt.Scan untuk array ikan. Lalu inputan dibaca oleh program menggunakan fmt.Scan dan &ikan[i], lalu nilai inputan dimasukkan ke array ikan berdasarkan index ([i]) secara berurutan.</li>
                </ol>
            </li>
            <li>Fungsi <strong>hitungWadah</strong>
                <br/>Fungsi ini adalah fungsi untuk menghitung total berat ikan di setiap wadah. Fungsi ini memiliki parameter pass by value yaitu ikan dengan tipe data arrIkan, x dan y dengan tipe data integer dan parameter pass by reference yaitu wadah dengan tipe data arrWadah. Parameter wadah diakses menggunakan simbol bintang (*) karena merupakan pointer ke tipe datanya. Tipe data yang dikembalikan adalah integer. Di dalam Fungsi ini terdapat:
                <ol>
                    <li>Membuat variabel baru menggunakan := dengan value (x + y - 1) dibagi y.</li>
                    <li>Perulangan for dengan inisiasi membuat variabel baru bernama w menggunakan := dan diberi value 0, kondisi variabel w lebih kecil dari variabel jumlahWadah, update w++ (Post-increment). Isi dari perulangan for ini adalah deklarasi variabel total dengan tipe data float64 dan diberi value 0.</li>
                    <li>Lalu di dalam perulangan tersebut terdapat perulangan for kedua, dengan inisiasi membuat variabel baru bernama i menggunakan := dan diberi value variabel w dikali y, kondisi variabel i lebih kecil dari variabel w ditambah 1 lalu dikali y dan (&&) variabel i lebih kecil dari variabel x, update i++ (Post-increment). Isi dari perulangan ini adalah variabel total dengan value variabel total ditambah elemen array ikan[i]</li>
                    <li>Pada baris akhir perulangan pertama, <strong>di luar perulangan kedua</strong>, terdapat array wadah[w] dengan value total</li>
                    <li>Return dari fungsi ini adalah variabel jumlahWadah</li>
                </ol>
            </li>
            <li>Fungsi <strong>rataWadah</strong>
                <br/>Fungsi ini adalah fungsi untuk menghitung total berat ikan dari setiap wadah. Fungsi ini memiliki parameter pass by value yaitu wadah dengan tipe data arrWadah dan jumlahWadah dengan tipe data integer. Tipe data return/kembalian dari fungsi ini adalah float64. Di dalam fungsi ini terdapat:
                <ol>
                    <li>variabel total dengan tipe data float64 dan diberi value 0.</li>
                    <li>Perulangan for dengan inisiasi membuat variabel baru bernama i menggunakan := dan diberi value 0, kondisi variabel i lebih kecil dari variabel jumlahWadah, update i++ (Post-increment).</li>
                    <li>Isi dari perulangan ini adalah variabel total dengan value variabel total ditambah elemen array wadah[i]</li>
                </ol>
                Return/kembalian dari fungsi ini adalah hasil dari variabel total dibagi konversi float64 dari variabel jumlahWadah.
            </li>
            <li>Prosedur <strong>beratPerWadah</strong>
                <br/>Prosedur ini adalah prosedur untuk menghitung menampilkan berat ikan per wadah. Prosedur ini memiliki parameter pass by reference yaitu jumlahWadah dengan tipe data integer, dan wadah dengan tipe data arrWadah. Kedua parameter tersebut diakses menggunakan simbol bintang (*) karena merupakan pointer ke tipe datanya. Di dalam prosedur ini terdapat:
                <ol>
                    <li>Teks menggunakan fmt.Print yang bertuliskan "Total berat per wadah: ".</li>
                    <li>Perulangan for dengan inisiasi membuat variabel baru bernama i menggunakan := dan diberi value 0, kondisi variabel i lebih kecil dari value variabel jumlahWadah, update i++ (Post-increment).</li>
                    <li>Isi dari perulangan ini adalah teks yang menampilkan total berat per wadah menggunakan fmt.Printf dan format %.2f (float dengan 2 angka di belakang koma) dari array wadah[i].</li>
                </ol>
            </li>
        </ol>
    </li>
</ul>

### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.

### Buatlah program dengan spesifikasi subprogram sebagai berikut:

#### type arrBalita [100]float64

#### func hitungMinMax(arrBerat arrBalita; bMin, bMax \*float64) {

#### /\* I.S. Terdefinisi array dinamis arrBerat

#### Proses: Menghitung berat minimum dan maksimum dalam array

#### F.S. Menampilkan berat minimum dan maksimum balita \*/

#### ... }

#### function rerata (arrBerat arrBalita) real {

#### /_ menghitung dan mengembalikan rerata berat balita dalam array _/

#### ... }

#### soal3.go

```go
package main

import "fmt"

const max = 1000001

type bilangan [max]int

func urutkan(angka *bilangan, n int) {
	var i, j, temp int
	for i = 1; i <= n-1; i+=1 {
		j = i
		temp = angka[j]
		for j > 0 && temp < angka[j-1] {
			angka[j] = angka[j-1]
			j = j - 1
		}
		angka[j] = temp
	}
}

func main() {
	var data bilangan
	var n, nilai int

	n = 0
	fmt.Println("Masukkan data angka bilangan bulat: ")
	fmt.Scan(&nilai)

	for nilai != -5313 {
		if nilai == 0 {
			urutkan(&data, n)
			if n%2 == 1 {
				fmt.Printf("Median dari %d data adalah: %d\n", n, data[n/2])
			} else {
				fmt.Printf("Median dari %d data adalah: %d\n", n, (data[n/2-1]+data[n/2])/2)
			}
		} else {
			data[n] = nilai
			n++
		}
		fmt.Scan(&nilai)
	}
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul14/output/output-soal3.png)

##### Penjelasan

Program ini adalah program untuk mencatat berat balita, tiap berat dimasukkan ke dalam array arrBerat dan menghitung berat minimum, maksimum dan rata-rata dari data tersebut. Berikut penjelasan yang lebih lanjut:

<ul>
    <li>Tipe array <strong>arrBalita</strong>
        <br/>Program memiliki tipe array bernama arrBalita dengan kapasitas 100 dan tipe datanya adalah float64.
    </li>
    <li><strong>Di dalam func main()</strong>
        <ol>
            <li>Program memiliki variabel balita dengan tipe data arrBalita, yang berarti balita adalah array untuk menyimpan banyaknya berat balita.</li>
            <li>Program memiliki variabel n dengan tipe data integer.</li>
            <li>Program memiliki variabel bMin (berat minimum/paling kecil) dan bMax (berat maksimum/paling besar) dengan tipe data float64.</li>
            <li>Program memanggil prosedur banyakBalita dengan parameter aktual &balita dan &n.</li>
            <li>Program memanggil prosedur hitungMinMax dengan argumen balita dan n, serta parameter aktual yaitu &bMin dan &bMax.</li>
            <li>Program memberi output menggunakan fmt.Printf (print format) untuk menampilkan variabel bMin</li>
            <li>Program memberi output menggunakan fmt.Printf (print format) untuk menampilkan variabel bMax</li>
            <li>Program memberi output menggunakan fmt.Printf (print format) untuk menampilkan hasil dari fungsi rerata dengan argumen balita dan n</li>
            <li>Variabel dan fungsi rerata pada ketiga output ditampilkan menggunakan format %.2f (format float dengan 2 angka di belakang koma)</li>
        </ol>
    </li>
    <li><strong>Prosedur dan fungsi</strong>
        <ol>
            <li>Prosedur <strong>banyakBalita</strong>
                <br/>Prosedur ini adalah prosedur untuk menerima input data jumlah balita dan juga memasukkan input data berat balita berdasarkan banyak balita. Prosedur ini memiliki parameter pass by reference yaitu arrBerat dengan tipe data arrBalita, serta n dengan tipe data integer. Kedua parameter diakses menggunakan simbol bintang (*) karena merupakan pointer ke tipe datanya. Di dalam prosedur ini terdapat:
                <ol>
                    <li>Teks menggunakan fmt.Print untuk menampilkan perintah memasukkan banyak data berat balita.</li>
                    <li>Input menggunakan fmt.Scan untuk variabel n, lalu hasil input dimasukkan ke variabel n.</li>
                    <li>Perulangan for dengan inisiasi membuat variabel baru bernama i menggunakan := dan diberi value 0, kondisi variabel i lebih kecil dari value variabel n, update i++ (Post-increment).</li>
                    <li>Isi dari perulangan ini adalah teks yang menampilkan index berat balita menggunakan fmt.Printf dan format %d (format integer) dari variabel i+1. Dan juga input menggunakan fmt.Scan untuk arrBerat. Lalu tiap hasil input dibaca oleh program menggunakan fmt.Scan dan &arrBerat[i], lalu nilai inputan dimasukkan ke arrBerat berdasarkan index ([i]) secara berurutan.</li>
                </ol>
            </li>
            <li>Prosedur <strong>hitungMinMax</strong>
                <br/>Prosedur ini adalah prosedur untuk menghitung elemen terkecil dan terbesar dari array arrBerat. prosedur ini memiliki parameter pass by value yaitu arrBerat dengan tipe data arrBalita, dan n dengan tipe data integer, serta parameter pass by reference yaitu bMin dan bMax dengan tipe data float64. Parameter pass by reference diakses menggunakan simbol bintang (*) karena merupakan pointer ke tipe datanya. Di dalam prosedur ini terdapat:
                <ol>
                    <li>Variabel j dengan tipe data integer dan diberi value 1.</li>
                    <li>Nilai yang ditunjuk pointer bMin diberi value array ke 0 dari arrBerat.</li>
                    <li>Nilai yang ditunjuk pointer bMax diberi value array ke 0 dari arrBerat.</li>
                    <li>Perulangan for dengan kondisi j lebih kecil dari n</li>
                    <li>Di dalam perulangan tersebut ada if pertama dengan kondisi jika ada elemen di arrBerat[j] yang lebih kecil dari value variabel bMin maka value dari variabel bMin adalah elemen arrBerat[j] yang paling kecil tersebut. Ada juga if kedua dengan kondisi jika ada elemen di arrBerat[j] yang lebih besar dari variabel bMax maka value dari value variabel bMax adalah elemen arrBerat[j] yang paling besar tersebut. Baris terakhir pada kode for ini adalah value dari variabel j adalah variabel j ditambah 1.</li>
                </ol>
            </li>
            <li>Fungsi <strong>rerata</strong>
                <br/>Fungsi ini adalah fungsi untuk menghitung rata-rata berat dari array data berat balita arrBerat. Fungsi ini memiliki parameter pass by value yaitu arrBerat dengan tipe data arrBalita dan n dengan tipe data integer. Tipe return/kembalian dari fungsi ini adalah float64. Di dalam fungsi ini terdapat:
                <ol>
                    <li>variabel total dengan tipe data float64 dan diberi value 0.</li>
                    <li>Perulangan for dengan inisiasi membuat variabel baru bernama i menggunakan := dan diberi value 0, kondisi variabel i lebih kecil dari variabel n, update i++ (Post-increment).</li>
                    <li>Isi dari perulangan ini adalah variabel total dengan value variabel total ditambah elemen array arrBerat[i]</li>
                </ol>
                Return/kembalian dari fungsi ini adalah hasil dari variabel total dibagi konversi float64 dari variabel n.
            </li>
        </ol>
    </li>
</ul>
