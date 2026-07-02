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
                    <li>Fmt.Scan untuk memasukkan banyaknya rumah di daerah ke-i, nomor tersebut dimasukkan ke variabel m.</li>
                    <li>Perulangan untuk memasukkan nomor rumah per daerah.</li>
                    <li>Program memanggil prosedur urutNomor dengan parameter aktual nomorRumah.</li>
                </ol>
            </li>
        </ol>
    </li>
    <li><strong>Prosedur urutNomor</strong>
        </br>Prosedur ini adalah prosedur selection sort untuk mengurutkan nomor rumah dari yang terkecil hingga yang terbesar. Berikut penjelasan yang lebih lanjut:
		<ol>
			<li>Prosedur ini memiliki pointer nomor dengan tipe data kerabat, dan n sebagai argumen bertipe data integer.</li>
			<li>Variabel t, i, j dan idx_min dengan tipe data integer.</li>
			<li>Perulangan for dengan inisiasi variabel i diberi value 1, kondisi variabel i lebih kecil sama dengan variabel n dikurangi 1, dan update i+=1 atau variabel i ditambahkan dengan variabel 1 di setiap perulangan. Perulangan ini berfungsi untuk mencari nilai terkecil di sisa array, lalu menukarnya ke posisi target (i-1), sehingga array nomor urut dari yang terkecil hingga terbesar.</li>
		</ol>
    </li>
</ul>

### 2. Belakangan diketahui ternyata Hercules itu tidak berani menyeberang jalan, maka selalu diusahakan agar hanya menyeberang jalan sesedikit mungkin, hanya diujung jalan. Karena nomor rumah sisi kiri jalan selalu ganjil dan sisi kanan jalan selalu genap, maka buatlah program kerabat dekat yang akan menampilkan nomor rumah mulai dari nomor yang ganjil lebih dulu terurut membesar dan kemudian menampilkan nomor rumah dengan nomor genap terurut mengecil.

### Format Masukan masih persis sama seperti sebelumnya.

### Keluaran terdiri dari n baris, yaitu rangkaian rumah kerabatnya terurut membesar untuk nomor ganjil, diikuti dengan terurut mengecil untuk nomor genap, di masing-masing daerah.

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

Program ini mengurutkan nomor rumah kerabat Hercules di setiap daerah dengan nomor ganjil diurutkan secara ascending dan nomor genap secara descending, menggunakan algoritma Selection Sort. Berikut penjelasan yang lebih lanjut:

<ul>
    <li>Tipe array <strong>rumahKerabat</strong>
    <br/>Program memiliki array bernama rumahKerabat dengan kapasitas rumah sebesar 1000000 dan tipe data integer.
    </li>
    <li><strong>Di dalam func main()</strong>
        <ol>
            <li>Program memiliki variabel n dan m dengan tipe data integer</li>
            <li>Program memiliki variabel ganjil dan genap dengan tipe data rumahKerabat, yang berarti kedua variabel tersebut adalah array.</li>
            <li>Program memiliki variabel i, j, k dan l dengan tipe data integer</li>
            <li>Program memiliki variabel nGanjil, nGenap dan val dengan tipe data integer</li>
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
                    <li>Fmt.Scan untuk memasukkan banyaknya rumah di daerah ke-i, nomor tersebut dimasukkan ke variabel m.</li>
                    <li>Variabel nGanjil dan nGenap yang diberi value 0.</li>
                    <li>Perulangan untuk memasukkan nomor rumah per daerah. Di dalam perulangan terdapat fmt.Scan sebagai input untuk variabel val, dan kondisi if else:
						<ol>
							<li>Kondisi 1, jika variabel val di-modulus 2 tidak sama dengan 0, maka angka yang diinputkan ke variabel val tersebut dimasukkan ke array ganjil dengan index nGanjil, dan variabel nGanjil bertambah di setiap perulangan (nGanjil++).</li>
							<li>Kondisi 2, jika kondisi pertama tidak terpenuhi maka akan menjalankan kondisi ini, atau angka yang dimasukkan ke variabel val adalah angka genap, angka yang diinputkan ke variabel val tersebut dimasukkan ke array genap dengan index nGenap, dan variabel nGenap bertambah di setiap perulangan (nGenap++).</li>
						</ol>
					</li>
                    <li>Program memanggil prosedur urutKecil dengan parameter aktual ganjil dan argumen nGanjil.</li>
                    <li>Program memanggil prosedur urutBesar dengan parameter aktual genap dan argumen nGenap.</li>
                    <li>Program menampilkan array yang sudah terurut menggunakan perulangan, masing-masing untuk array ganjil dan genap.</li>
                </ol>
            </li>
        </ol>
    </li>
    <li><strong>Prosedur urutKecil</strong>
        </br>Prosedur ini adalah prosedur selection sort untuk mengurutkan nomor rumah dari yang terkecil hingga yang terbesar. Berikut penjelasan yang lebih lanjut:
		<ol>
			<li>Prosedur ini memiliki pointer nomor dengan tipe data rumahKerabat, dan n sebagai parameter pass by value bertipe data integer.</li>
			<li>Variabel t, i, j dan idx_min dengan tipe data integer.</li>
			<li>Perulangan for dengan inisiasi variabel i diberi value 1, kondisi variabel i lebih kecil sama dengan variabel n dikurangi 1, dan update i+=1 atau variabel i ditambahkan dengan variabel 1 di setiap perulangan. Perulangan ini berfungsi untuk mencari nilai terkecil di sisa array, lalu menukarnya ke posisi target (i-1), sehingga array nomor urut dari yang terkecil hingga terbesar.</li>
		</ol>
    </li>
    <li><strong>Prosedur urutBesar</strong>
        </br>Prosedur ini adalah prosedur selection sort untuk mengurutkan nomor rumah dari yang terbesar hingga yang terkecil. Berikut penjelasan yang lebih lanjut:
		<ol>
			<li>Prosedur ini memiliki pointer nomor dengan tipe data rumahKerabat, dan n sebagai parameter pass by value bertipe data integer.</li>
			<li>Variabel t, i, j dan idx_max dengan tipe data integer.</li>
			<li>Perulangan for dengan inisiasi variabel i diberi value 1, kondisi variabel i lebih kecil sama dengan variabel n dikurangi 1, dan update i+=1 atau variabel i ditambahkan dengan variabel 1 di setiap perulangan. Perulangan ini berfungsi untuk mencari nilai terbesar di sisa array, lalu menukarnya ke posisi target (i-1), sehingga array nomor urut dari yang terbesar hingga terkecil.</li>
		</ol>
    </li>
</ul>

### 3. Kompetisi pemrograman yang baru saja berlalu diikuti oleh 17 tim dari berbagai perguruan tinggi ternama. Dalam kompetisi tersebut, setiap tim berlomba untuk menyelesaikan sebanyak mungkin problem yang diberikan. Dari 13 problem yang diberikan, ada satu problem yang menarik. Problem tersebut mudah dipahami, hampir semua tim mencoba untuk menyelesaikannya, tetapi hanya 3 tim yang berhasil. Apa sih problemnya?

### "Median adalah nilai tengah dari suatu koleksi data yang sudah terurut. Jika jumlah data genap, maka nilai median adalah rerata dari kedua nilai tengahnya. Pada problem ini, semua data merupakan bilangan bulat positif, dan karenanya rerata nilai tengah dibulatkan ke bawah."

### Buatlah program median yang mencetak nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0.

### Masukan berbentuk rangkaian bilangan bulat. Masukan tidak akan berisi lebih dari 1000000 data, tidak termasuk bilangan 0. Data 0 merupakan tanda bahwa median harus dicetak, tidak termasuk data yang dicari mediannya. Data masukan diakhiri dengan bilangan bulat -5313.

### Keluaran adalah median yang diminta, satu data per baris.

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
Program ini mencari nilai median terhadap seluruh data yang sudah terbaca, jika data yang dibaca saat itu adalah 0. Berikut penjelasan yang lebih lanjut:

<ul>
    <li>Tipe array <strong>bilangan</strong>
    <br/>Program memiliki array bernama bilangan dengan kapasitas max sebesar 1000001 dan tipe data integer.
    </li>
    <li><strong>Di dalam func main()</strong>
        <ol>
            <li>Program memiliki variabel data dengan tipe data bilangan, yang berarti variabel data adalah sebuah array.</li>
            <li>Program memiliki variabel n dan nilai dengan tipe data integer.</li>
            <li>Program memberi value 0 ke variabel n.</li>
            <li>User diberi perintah untuk menginputkan variabel nilai menggunakan fmt.Scan, perintah tersebut ditulis menggunakan fmt.Println.</li>
            <li>Program membuat perulangan for, dengan kondisi nilai tidak sama dengan -5313, yang berarti perulangan akan tetap berjalan sampai variabel nilai memiliki value -5313</li>
            <li>Kode di dalam perulangan for:
                <ol type="a">
                    <li>Kondisi if, jika nilai sama dengan 0, maka program akan menjalankan prosedur urutkan dengan parameter aktual yaitu variabel data dan argumen yaitu variabel n. Terdapat kondisi if lagi, jika variabel n di-modulus 2 hasilnya sama dengan 1, maka program akan menampilkan median dari array data yang indexnya dibagi 2. Jika kondisi tersebut tidak terpenuhi maka median yang ditampilkan adalah array data dengan index n dibagi 2-1, ditambah array data dengan index n dibagi 2. Setelah ditemukan hasilnya, array tersebut dibagi 2 lagi.</li>
                    <li>Kondisi else, yang berarti kondisi sebelumnya tidak terpenuhi, isi dari kondisi ini adalah array data dengan index n diisi variabel nilai, variabel n bertambah di setiap perulangan (n++).</li>
                    <li>Terakhir ada fmt.Scan sebagai input untuk variabel nilai.</li>
                </ol>
            </li>
        </ol>
    </li>
    <li><strong>Prosedur urutkan</strong>
        </br>Prosedur ini adalah prosedur untuk mengurutkan angka array dari yang terkecil hingga yang terbesar. Berikut penjelasan yang lebih lanjut:
		<ol>
			<li>Prosedur ini memiliki pointer angka dengan tipe data bilangan, dan n sebagai argumen bertipe data integer.</li>
			<li>Variabel i, j dan temp dengan tipe data integer.</li>
			<li>Perulangan for dengan inisiasi variabel i diberi value 1, kondisi variabel i lebih kecil sama dengan variabel n dikurangi 1, dan update i+=1 atau variabel i ditambahkan dengan variabel 1 di setiap perulangan. Perulangan ini berfungsi untuk mengurutakan angka dengan cara elemen dengan index ke-i disimpan sementara di variabel temp, lalu digeser mundur (angka[j] = angka[j-1]) selama elemen sebelumnya lebih besar, sampai ditemukan posisi yang tepat.</li>
		</ol>
    </li>
</ul>