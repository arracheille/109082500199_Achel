# <h1 align="center">Laporan Praktikum Modul 14 - Soal Latihan Insertion Sort</h1>

<p align="center">Aqilla Rachel Rabbani - 109082500199</p>

## Unguided

### 1. Buatlah sebuah program yang digunakan untuk membaca data integer seperti contoh yang diberikan di bawah ini, kemudian diurutkan (menggunakan metoda insertion sort), dan memeriksa apakah data yang terurut berjarak sama terhadap data sebelumnya.

### Masukan terdiri dari sekumpulan bilangan bulat yang diakhiri oleh bilangan negatif. Hanya bilangan non negatif saja yang disimpan ke dalam array.

### Keluaran terdiri dari dua baris. Baris pertama adalah isi dari array setelah dilakukan pengurutan, sedangkan baris kedua adalah status jarak setiap bilangan yang ada di dalam array. "Data berjarak x" atau "data berjarak tidak tetap".

#### soal1.go

```go
package main

import "fmt"

type arrInt [1000001]int

func insertionSort1(T *arrInt, n int) {
	var temp, i, j int

	for i = 1; i <= n-1; i+=1 {
		j = i
		temp = T[j]
		for j > 0 && temp < T[j-1] {
			T[j] = T[j-1]
			j = j - 1
		}
		T[j] = temp
	}
}

func main() {
	var T arrInt
	var n, x int
	var selesai bool

	n = 0
	selesai = false

	fmt.Print("Masukkan data nomor: ")
    for !selesai {
        fmt.Scan(&x)

        if x < 0 {
            selesai = true
        } else {
            T[n] = x
            n = n + 1
        }
    }

	insertionSort1(&T, n)

	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(T[i])
	}
	fmt.Println()

	if n <= 1 {
		fmt.Println("Data berjarak tidak tetap")
	} else {
		jarak := T[1] - T[0]
		tetap := true
		for j := 2; j < n; j+=1 {
			if T[j]-T[j-1] != jarak {
				tetap = false
			}
		}
		if tetap {
			fmt.Println("Data berjarak", jarak)
		} else {
			fmt.Println("Data berjarak tidak tetap")
		}
	}
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul14-2/output/output-soal1.png)

##### Penjelasan

Program ini adalah sebuah program untuk mengurutkan angka berdasarkan insertionsort. Berikut penjelasan yang lebih lanjut:

<ul>
  <li>Program memiliki tipe data (array) arrInt dengan kapasitas 1000001, bertipe data integer.</li>
  <li>Di dalam fungsi main, program memiliki variabel T dengan tipe data arrInt, yang berarti variabel T adalah array. Variabel n dan x dengan tipe data integer, variabel selesai dengan tipe data boolean.</li>
  <li>Variabel n diberi nilai 0 dan variabel selesai diberi nilai false.</li>
  <li>Program membuat perulangan repeat-until, isi di dalam perulangan adalah input nilai variabel x, setelah memasukkan input, inputan user dibaca oleh program menggunakan Scan. Di dalam perulangan juga terdapat kondisi if-else:
	<ul>
		<li>Kondisi 1 (if), jika x kurang dari 0 maka variabel selesai bernilai true, perulangan berhenti.</li>
		<li>Kondisi 2 (else), jika kondisi pertama tidak dipenuhi maka akan menjalankan kondisi ini, yaitu variabel x dimasukkan ke array T dengan index ke-n, dan variabel n ditambah 1 di setiap perulangan</li>
	</ul>
  </li>
  <li>Program menjalankan prosedur insertionSort1 dengan pointer variabel T dan argumen variabel n. Isi dari prosedur tersebut adalah:
	<ul>
		<li>Variabel temp, i dan j dengan tipe data integer</li>
		<li>Perulangan untuk mencari nilai terkecil dari array T dan mengurutkannya secara ascending berdasarkan index ke-j (variabel j diberi nilai variabel i). Pengurutan dilakukan dengan cara menukar elemen array T dengan array variabel temp.</li>
	</ul>
  </li>
  <li>Perulangan untuk menampilkan array T urut dari index ke-0 sampai index ke-n, jika i lebih dari 0 maka yang ditampilkan adalah spasi.</li>
  <li>Program menjalankan fmt.Println agar output selanjutnya berada di baris baru.</li>
  <li>Program memiliki kondisi if-else:
	<ul>
		<li>Kondisi 1 (if), jika n kurang dari sama dengan 1 maka program menampilkan teks "Data berjarak tidak tetap".</li>
		<li>Kondisi 2 (else), jika kondisi pertama tidak dipenuhi maka akan menjalankan kondisi ini, dengan tujuan untuk mencari jarak antar 2 elemen array yang sudah diurutkan. Kondisi ini membuat variabel jarak dengan nilai array T dengan index ke-1 dikurangi index ke-0, dan juga variabel tetap bernilai true. Terdapat juga perulangan untuk mengecek setiap elemen array dan mengecek apakah jarak antar elemen array sama dengan nilai variabel jarak. Jika jarak data tidak tetap maka variabel tetap bernilai false. Terdapat juga kondisi if-else:
			<ul>
				<li>Kondisi 1 (if), jika variabel tetap bernilai true, maka program menampilkan teks "Data berjarak (isi nilai variabel jarak)" menggunakan fmt.Println.</li>
				<li>Kondisi 2 (else), jika kondisi pertama tidak dipenuhi maka akan menjalankan kondisi ini, maka program menampilkan teks "Data berjarak tidak tetap" menggunakan fmt.Println.</li>
			</ul>
		</li>
	</ul>
  </li>
</ul>

### 2. Sebuah program perpustakaan digunakan untuk mengelola data buku di dalam suatu perpustakaan.

### Masukan terdiri dari beberapa baris. Baris pertama adalah bilangan bulat N yang menyatakan banyaknya data buku yang ada di dalam perpustakaan. N baris berikutnya, masing-masingnya adalah data buku sesuai dengan atribut atau field pada struct. Baris terakhir adalah bilangan bulat yang menyatakan rating buku yang akan dicari.

### Keluaran terdiri dari beberapa baris. Baris pertama adalah data buku terfavorit, baris kedua adalah lima judul buku dengan rating tertinggi, selanjutnya baris terakhir adalah data buku yang dicari sesuai rating yang diberikan pada masukan baris terakhir.

#### soal2.go

```go
package main

import "fmt"

const nMax = 7919

type Buku struct {
    id, judul, penulis, penerbit string
    eksemplar, tahun, rating int
}

type DaftarBuku [nMax + 1]Buku

var Pustaka DaftarBuku
var nPustaka int

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
    var b Buku
	fmt.Println("Masukkan jumlah data buku: ")
    fmt.Scan(n)
	fmt.Println("\nMasukkan data buku:")
    for i := 1; i <= *n; i++ {
        fmt.Scan(&b.id, &b.judul, &b.penulis, &b.penerbit, &b.eksemplar, &b.tahun, &b.rating)
        pustaka[i] = b
    }
}

func CetakTerfavorit(pustaka *DaftarBuku, n int) {
	fmt.Println("\nBuku ter-favorit (urut dari rating tertinggi):")
    idxMax := 1
    for j := 2; j <= n; j++ {
        if pustaka[j].rating > pustaka[idxMax].rating {
            idxMax = j
        }
    }
    b := pustaka[idxMax]
    fmt.Println(b.judul, b.penulis, b.penerbit, b.tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
    var temp Buku
    var k, l int

    for k = 2; k <= n; k++ {
        temp = pustaka[k]
        l = k
        for l > 1 && temp.rating > pustaka[l-1].rating {
            pustaka[l] = pustaka[l-1]
            l = l - 1
        }
        pustaka[l] = temp
    }
}

func Cetak5Terbaru(pustaka *DaftarBuku, n int) {
	fmt.Println("\n5 Buku dengan rating tertinggi:")
    batas := 5
    if n < batas {
        batas = n
    }
    for m := 1; m <= batas; m++ {
        fmt.Println(pustaka[m].judul)
    }
}

func CariBuku(pustaka *DaftarBuku, n int, r int) {
    low, high := 1, n
    ketemu := -1
    for low <= high {
        mid := (low + high) / 2
        if pustaka[mid].rating == r {
            ketemu = mid
            break
        } else if pustaka[mid].rating < r {
            high = mid - 1
        } else {
            low = mid + 1
        }
    }

    if ketemu == -1 {
        fmt.Println("Tidak ada buku dengan rating seperti itu")
    } else {
        b := pustaka[ketemu]
        fmt.Println(b.judul, b.penulis, b.penerbit, b.tahun, b.eksemplar, b.rating)
    }
}

func main() {
    var n, r int

    DaftarkanBuku(&Pustaka, &n)
    CetakTerfavorit(&Pustaka, n)
    UrutBuku(&Pustaka, n)
    Cetak5Terbaru(&Pustaka, n)

    fmt.Print("\nMasukkan rating yang akan dicari: ")
    fmt.Scan(&r)
    CariBuku(&Pustaka, n, r)
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_2](https://github.com/arracheille/109082500199_Achel/blob/main/Modul14-2/output/output-soal2.png)

##### Penjelasan

Program ini adalah program untuk mendata data buku, menampilkan buku ter-favorit dan 5 buku dengan rating tertinggi. Pengguna bsia mencari buku berdasarkan rating yang akan di cari di akhir program. Berikut penjelasan yang lebih lanjut:

<ul>
  <li>Program memiliki const/konstanta dengan nama nMax bernilai 7919.</li>
  <li>Program memiliki tipe struct dengan nama Buku dengan isi:
	<ul>
		<li>Variabel id, judul, penulis dan penerbit dengan tipe data string</li>
		<li>Variabel eksemplar, tahun dan rating dengan tipe data integer</li>
	</ul>
  </li>
  <li>Program memiliki tipe data (array) DaftarBuku dengan kapasitas nilai variabel nMax ditambah 1, bertipe data Buku. Ini berarti setiap elemen di dalam array tersebut terdapat data buku.</li>
  <li>Program memiliki variabel Pustaka dengan tipe data DaftarBuku, berarti variabel ini adalah array. Dan juga variabel nPustaka dengan tipe data integer.</li>
  <li><strong>Di dalam function main:</strong>
	<ul>
		<li>Terdapat variabel n dan r dengan tipe data integer.</li>
		<li>Program menjalankan prosedur DaftarkanBuku (Variabel pustaka dan n sebagai pointer), CetakTerfavorit, UrutBuku, dan Cetak5Terbaru. Pada ketiga prosedur terakhir variabel Pustaka sebagai pointer, variabel n sebagai argumen.</li>
		<li>Terdapat input nilai variabel r, setelah memasukkan input, inputan user dibaca oleh program menggunakan Scan.</li>
		<li>Program menjalankan prosedur CariBuku (variabel Pustaka sebagai pointer, variabel n dan r sebagai argumen)</li>
	</ul>
  </li>
  <li><strong>Prosedur DaftarkanBuku:</strong>
	<ul>
		<li>Terdapat variabel b dengan tipe data Buku.</li>
		<li>Terdapat input nilai variabel n, setelah memasukkan input, inputan user dibaca oleh program menggunakan Scan.</li>
		<li>Perulangan untuk memasukkan data buku sebanyak nilai variabel n, di dalamnya terdapat input nilai variabel b berdasarkan struct, setelah memasukkan input, inputan user dibaca oleh program menggunakan Scan. Setiap data yang dimasukkan oleh user dimasukkan ke array pustaka berdasarkan index ke-i, index i yang awalnya 1 bertambah 1 di setiap perulangan.</li>
	</ul>
  </li>
  <li><strong>Prosedur CetakTerfavorit:</strong>
	<ul>
		<li>Terdapat variabel idxMax dengan nilai 1</li>
		<li>Perulangan untuk mencari buku dengan rating tertinggi.</li>
		<li>Setelah perulangan selesai dijalankan, terdapat variabel b dengan nilai array pustaka dengan index ke-idxMax</li>
		<li>Teks output yang berisi judul, penulis, penerbit dan tahun keluaran buku.</li>
	</ul>
  </li>
  <li><strong>Prosedur UrutBuku:</strong>
	<ul>
		<li>Terdapat variabel temp dengan tipe data Buku, dan variabel k dan l dengan tipe data integer.</li>
		<li>Perulangan untuk mengurutkan buku berdasarkan rating tertinggi secara descending (tertinggi ke terendah). Terdapat variabel temp yang diberi nilai array pustaka dengan index ke-k. Pengurutan dilakukan dengan cara menggeser elemen-elemen sebelumnya yang memiliki rating lebih kecil akan digeser ke kanan. Nilai temp dimasukkan sesudai dengan index ke-l</li>
		<li>Setelah perulangan selesai dijalankan, terdapat array pustaka dengan index ke-l dan elemen di dalam array adalah variabel temp.</li>
	</ul>
  </li>
  <li><strong>Prosedur Cetak5Terbaru:</strong>
	<br />
	Prosedur ini berfungsi untuk menampilkan data buku yang sudah diurutkan oleh prosedur UrutBuku, dan menampilkan 5 judul data buku dengan rating tertinggi.
  </li>
  <li><strong>Prosedur CariBuku:</strong>
	<ul>
		<li>Terdapat variabel low dengan nilai 1, variabel high dengan nilai n, variabel ketemu dengan nilai -1 (jika -1 maka buku belum ditemukan)</li>
		<li>Perulangan untuk mencari buku berdasarkan input user (variabel r). Pencarian yang digunakan pada perulangan ini adalah binary search. Jika program menemukan elemen array yang sama dengan inputan user, maka nilai variabel ketemu adalah index dari elemen tersebut.</li>
		<li>Terdapat kondisi if-else:
			<ul>
				<li>Kondisi 1 (if), Jika nilai variabel ketemu adalah -1 (tidak ada buku dengan rating yang diinputkan user) maka program akan menampilkan teks "Tidak ada buku dengan rating seperti itu".</li>
				<li>Kondisi 2 (else), jika kondisi pertama tidak dipenuhi maka akan menjalankan kondisi ini, yaitu variabel b diisi dengan array pustaka dengan index ke-ketemu, lalu menampilkan data buku berdasarkan variabel b.</li>
			</ul>
	</ul>
  </li>
</ul>
