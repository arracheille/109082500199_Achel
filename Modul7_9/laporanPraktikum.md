# <h1 align="center">Laporan Praktikum Modul 9 - Soal Latihan Modul 7 dan 9 </h1>

<p align="center">Aqilla Rachel Rabbani - 109082500199</p>

## Unguided

### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.
### Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat.
### Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".
### Fungsi untuk menghitung jarak titik (a, b) dan (c, d) dimana rumus jarak adalah:
### jarak = √(a − c)^2 + (b - d)^2
### dan juga fungsi untuk menentukan posisi sebuah titik sembarang berada di dalam suatu lingkaran atau tidak.
#### function jarak(p, q : titik) -> real
#### {Mengembalikan jarak antara titik p(x,y) dan titik q(x,y)}
#### function didalam(c:lingkaran, p:titik) -> boolean
#### {Mengembalikan true apabila titik p(x,y) berada di dalam lingkaran c yang memiliki titik pusat (cx,cy) dan radius r}
### Catatan: Lihat paket math dalam lampiran untuk menggunakan fungsi math.Sqrt() untukmenghitung akar kuadrat.

#### soal1.go

```go
package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	pusat titik
	r int
}

func jarak(p, q titik) float64 {
	dx := float64(p.x - q.x)
	dy := float64(p.y - q.y)
	return math.Sqrt(dx*dx + dy*dy)
}

func didalam(c lingkaran, p titik) bool {
	return jarak(c.pusat, p) <= float64(c.r)
}

func main() {
	var (
		c1, c2 lingkaran
		p titik
		output string
	)

	fmt.Print("Masukkan koordinat titik pusat x, y dan radius dari lingkaran 1: ")
	fmt.Scanln(&c1.pusat.x, &c1.pusat.y, &c1.r)
	fmt.Print("Masukkan koordinat titik pusat x, y dan radius dari lingkaran 2: ")
	fmt.Scanln(&c2.pusat.x, &c2.pusat.y, &c2.r)
	fmt.Print("Masukkan koordinat titik sembarang (x, y): ")
	fmt.Scan(&p.x, &p.y)

	titik1 := didalam(c1, p)
	titik2 := didalam(c2, p)

	if titik1 && titik2 {
		output = "Titik di dalam lingkaran 1 dan 2"
	} else if titik1 {
		output = "Titik di dalam lingkaran 1"
	} else if titik2 {
		output = "Titik di dalam lingkaran 2"
	} else {
		output = "Titik di luar lingkaran 1 dan 2"
	}

	fmt.Print(output)
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul7_9/output/soal1.png)

##### Penjelasan
Program ini adalah program untuk menentukan posisi sebuah titik sembarang (variabel x dan variabel y) berdasarkan dua lingkaran. Berikut penjelasan yang lebih lanjut:
<ul>
	<li><strong>Struct</strong>
		<ol>
			<li><strong>titik</strong>
			</br>Struct titik memiliki variabel x dan y dengan tipe data integer.
			</li>
			<li><strong>lingkaran</strong>
			</br>Struct lingkaran memiliki variabel pusat dengan tipe data titik, dan variabel r dengan tipe data integer.
			</li>
		</ol>
	</li>
	<li><strong>Di dalam func main()</strong>
		<ol>
			<li>Program memiliki variabel c1 (lingkaran 1) dan c2 (lingkaran 2) dengan tipe data <strong>lingkaran</strong>, variabel p dengan tipe data <strong>titik</strong> dan variabel output dengan tipe data string</li>
			<li>Program menampilkan text perintah untuk memasukkan nilai semua variabel kecuali variabel output kepada user menggunakan fmt.Print.</li>
			<li>Program memberikan input untuk nilai seluruh variabel kecuali variabel output. Setelah memasukkan input, inputan user dibaca oleh program menggunakan fmt.Scan dan &(variabel), lalu nilai inputan dimasukkan ke semua variabel kecuali variabel output secara berurutan. Input untuk variabel dari lingkaran 1 dan lingkaran 2 menggunakan fmt.Scanln agar kode selanjutnya berada di line/baris baru.</li>
			<li>Text perintah dan fmt.Scan dituliskan secara berkelompok, yang pertama untuk lingkaran 1, lalu lingkaran 2 dan titik sembarang.</li>
			<li><strong>Contoh cara kerja struct</strong>
			</br>Di dalam fmt.Scan, program memanggil variabel menggunakan struct. Contohnya c1.pusat.x, karena pada deklarasi variabel c1 tipe datanya lingkaran, maka variabel c1 berelasi dengan variabel dari struct lingkaran yaitu variabel pusat dan r. Lalu, variabel pusat bertipe data titik, yang berarti variabel pusat memiliki relasi dengan variabel dari struct titik yaitu x dan y. Jadi, c1.pusat.x berarti program memanggil variabel x dari variabel pusat dari variabel c1, semua field-nya saling berelasi membentuk hierarki dari c1 → pusat → x.</li>
			<li>Program membuat variabel bernama titik1 dengan isi variabel adalah function didalam, variabel ini menghitung function didalam dari nilai variabel c1 dan p.</li>
			<li>Program membuat variabel bernama titik2 dengan isi variabel adalah function didalam, variabel ini menghitung function didalam dari nilai variabel c2 dan p.</li>
		</ol>
	</li>
    <li><strong>Fungsi/function(func) didalam dan jarak</strong>
		<ol>
			<li><strong>didalam</strong>
			</br>Pada function ini, program memiliki parameter c dengan tipe data lingkaran dan p dengan tipe data titik, return type dari function adalah boolean. Nilai variabel yang dimasukkan ke function dari func main() dimasukkan ke variabel c dan p secara berurutan. Function ini mengecek apakah jarak antara titik dan pusat lingkaran lebih kecil sama dengan radius dengan cara menggunakan function jarak yang berisi titik pusat lingkaran (c.pusat) dan koordinat titik sembarang (p), lalu dicek menggunakan <= dan konversi float64 dari variabel radius dari lingkaran (c.r). Kode tersebut dikembalikan menggunakan return.
			</li>
			<li><strong>jarak</strong>
			</br>Pada function ini, program menggunakan parameter p dan q dengan tipe data titik, return type dari function adalah float64. Nilai variabel c.pusat dan p dari function didalam dimasukkan ke variabel p dan q secara berurutan. <strong>Berarti p.x = variabel x (struct titik) dari variabel c.pusat, p.y = variabel y (struct titik) dari variabel c.pusat, q.x = variabel x (struct titik) dari variabel p, q.y = variabel y (struct titik) dari variabel p</strong>. Function ini membuat variabel baru bernama dx menggunakan := untuk menghitung hasil pengurangan dari variabel x dari variabel p dan variabel x dari variabel q, dan juga variabel baru bernama dy menggunakan := untuk menghitung hasil pengurangan dari variabel y dari variabel p dan variabel y dari variabel q, kedua hasil dikonversikan ke float64 di dalam masing-masing variabel baru. Lalu program mengembalikan nilai dari jarak yaitu akar (math.Sqrt) dari dx dikali dx dan dy dikali dy.
			</li>
		</ol>
	</li>
	<li>Setelah menghitung titik1 dan titik2, program menggunakan if-else-if untuk membandingkan kondisi:
		<ol>
			<li><strong>Kondisi 1 (if)</strong>
			</br>Kondisinya adalah jika titik1 dan (&&) titik2 bernilai true, maka variabel output berisi "Titik di dalam lingkaran 1 dan 2"</li>
			<li><strong>Kondisi 2 (else if)</strong>
			</br>Kondisinya adalah jika hanya titik1 yang bernilai true, maka variabel output berisi "Titik di dalam lingkaran 1"</li>
			<li><strong>Kondisi 3 (else if)</strong>
			</br>Kondisinya adalah jika hanya titik2 yang bernilai true, maka variabel output berisi "Titik di dalam lingkaran 2"</li>
			<li><strong>Kondisi 4 (else)</strong>
			</br>Kondisinya adalah jika semua kondisi diatas tidak terpenuhi, atau berarti titik1 dan titik2 bernilai false, maka variabel output berisi "Titik di luar lingkaran 1 dan 2".</li>
		</ol>
	</li>
	<li>Terakhir, ada fmt.Print untuk menampilkan keluaran/text dari nilai variabel output</li>
</ul>

### 2. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu. Program dapat menampilkan beberapa informasi berikut:
### a. Menampilkan keseluruhan isi dari array.
### b. Menampilkan elemen-elemen array dengan indeks ganjil saja.
### c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah genap).
### d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna.
### e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid. Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil
### f. Menampilkan rata-rata dari bilangan yang ada di dalam array.
### g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut.
### h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut.

#### soal2.go

```go
package main

import (
	"fmt"
	"math"
)

const array_max int = 100

type tabelInt [array_max]int

func array(t *tabelInt, n *int) {
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Printf("Elemen ke-%d: ", i)
		fmt.Scan(&t[i])
	}
}

func tampilSemua(t tabelInt, n int) {
	fmt.Print("Semua elemen: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", t[i])
	}
	fmt.Println()
}

func tampilGanjil(t tabelInt, n int) {
	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Printf("%d ", t[i])
	}
	fmt.Println()
}

func tampilGenap(t tabelInt, n int) {
	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Printf("%d ", t[i])
	}
	fmt.Println()
}

func tampilKelipatanX(t tabelInt, n, x int) {
	fmt.Printf("Indeks kelipatan %d: ", x)
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Printf("%d ", t[i])
		}
	}
	fmt.Println()
}

func hapusElemen(t *tabelInt, n *int, idx int) {
	for i := idx; i < *n-1; i++ {
		t[i] = t[i+1]
	}
	*n--
}

func rataRata(t tabelInt, n int) float64 {
	sum := 0
	for i := 0; i < n; i++ {
		sum += t[i]
	}
	return float64(sum) / float64(n)
}

func stdDeviasi(t tabelInt, n int) float64 {
	mean := rataRata(t, n)
	sum := 0.0
	for i := 0; i < n; i++ {
		diff := float64(t[i]) - mean
		sum += diff * diff
	}
	return math.Sqrt(sum / float64(n))
}

func frekuensi(t tabelInt, n, cari int) int {
	freq := 0
	for i := 0; i < n; i++ {
		if t[i] == cari {
			freq++
		}
	}
	return freq
}

func main() {
	var tab tabelInt
	var n int

	array(&tab, &n)

	tampilSemua(tab, n)

	tampilGanjil(tab, n)

	tampilGenap(tab, n)

	var x int
	fmt.Print("Masukkan nilai x untuk kelipatan: ")
	fmt.Scan(&x)
	tampilKelipatanX(tab, n, x)

	var idx int
	fmt.Print("Masukkan indeks yang akan dihapus: ")
	fmt.Scan(&idx)
	hapusElemen(&tab, &n, idx)
	fmt.Print("Array setelah penghapusan: ")
	tampilSemua(tab, n)

	fmt.Printf("Rata-rata: %.2f\n", rataRata(tab, n))

	fmt.Printf("Standar deviasi: %.2f\n", stdDeviasi(tab, n))

	var cari int
	fmt.Print("Cari frekuensi bilangan: ")
	fmt.Scan(&cari)
	fmt.Printf("Frekuensi %d: %d\n", cari, frekuensi(tab, n, cari))
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul7_9/output/soal2.png)

##### Penjelasan

### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga.
### Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja.
### Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.

#### soal3.go

```go
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
		fmt.Printf("Hasil %d : %s", i+1, hasil[i])
	}
	fmt.Println("Pertandingan selesai")
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul7_9/output/soal3.png)

##### Penjelasan

Program ini adalah program untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Berikut penjelasan yang lebih lanjut:

<ul>
  	<li><strong>Tipe data hasil_skor</strong>
	</br>hasil_skor adalah tipe data array dengan kapasitas hasil_max (100) elemen bertipe string. Kapasitas hasil_max bertipe data int dengan nilai 100, karena menggunakan const, berarti kapasitasnya tidak bisa berubah (immutable).
	</li>
  	<li><strong>Di dalam func main()</strong>
		<ol>
			<li>Program memiliki variabel klub_a dan klub_b dengan tipe data string, variabel skor_a dan skor_b dengan tipe data integer, variabel hasil dengan tipe data hasil_skor dan variabel n dengan tipe data integer</li>
			<li>Program menampilkan text perintah untuk memasukkan nilai variabel klub_a dan klub_b kepada user menggunakan fmt.Print.</li>
			<li>Program memberikan input untuk nilai variabel klub_a dan klub_b. Setelah memasukkan input, inputan user dibaca oleh program menggunakan fmt.Scan dan &klub_a dan &klub_b, lalu nilai inputan dimasukkan ke variabel klub_a dan klub_b.</li>
			<li>Teks perintah dan input dibagi menjadi 2 kelompok yaitu klub a dan klub b.</li>
			<li>Program membuat variabel baru bernama pertandingan menggunakan := dan diberi nilai 1</li>
			<li>Program membuat perulangan for tanpa inisiasi, di dalam perulangan for terdapat:
				<ol type="a">
					<li>Text perintah untuk memasukkan nilai variabelskorb_a dan skor_b kepada user menggunakan fmt.Printf (print format). Print ini berisi variabel pertandingan</li>
					<li>Input untuk nilai variabel skor_a dan skor_b. Setelah memasukkan input, inputan user dibaca oleh program menggunakan fmt.Scan dan &skor_a, &klub_b, lalu nilai inputan dimasukkan ke variabel skor_a dan klub_b secara berurutan.</li>
					<li>Kondisi if, dimana jika variabel skor_a kurang dari 0 atau skor_b kurang dari 0 maka perulangan akan berhenti</li>
					<li>Kondisi if dengan kondisi 1: jika skor a lebih besar dari skor b, maka variabel hasil berisi klub_a, kondisi 2: jika skor b lebih besar dari skor a maka variabel hasil berisi klub_b, kondisi 3: jika kedua kondisi sebelumnya tidak terpenuhi maka variabel hasil berisi teks "Draw".</li>
					<li>Pada bagian akhir for dari n++ dan pertandingan++ yang berarti pada setiap perulangan nilai variabel n dan pertandingan bertambah 1</li>
				</ol>
			</li>
			<li>Program membuat perulangan for dengan:
				<ol type="a">
					<li><strong>inisiasi</strong>, yaitu program membuat variabel baru bernama i menggunakan “:=” dan diberi nilai 0.</li>
					<li><strong>kondisinya</strong> adalah nilai variabel i kurang dari n.</li>
					<li><strong>updatenya</strong> adalah i++ (Post-increment).</li>
				</ol>
				</br>Kode di dalam perulangan ini berisi output hasil dan iterasinya.
			</li>
		</ol>
	</li>
</ul>

### 4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom.
### Lengkapi potongan algoritma berikut ini!

```go
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune
	tab : tabel
	m : integer

func isiArray(t *tabel, n *int)
/*I.S. Data tersedia dalam piranti masukan
F.S. Array t berisi sejumlah n karakter yang dimasukkan user,
Proses input selama karakter bukanlah TITIK dan n <= NMAX */

func cetakArray(t tabel, n int)
/*I.S. Terdefinisi array t yang berisi sejumlah n karakter
F.S. n karakter dalam array muncul di layar */

func balikanArray(t *tabel, n int)
/*I.S. Terdefinisi array t yang berisi sejumlah n karakter
F.S. Urutan isi array t terbalik */

func main(){
	var tab tabel
	var m int
	// si array tab dengan memanggil prosedur isiArray
	// Balikian isi array tab dengan memanggil balikanArray
	// Cetak is array tab
}
```

### Modifikasi program tersebut dengan menambahkan fungsi palindrom. Tambahkan instruksi untuk memanggil fungsi tersebut dan menampilkan hasilnya pada program utama.
### *Palindrom adalah teks yang dibaca dari awal atau akhir adalah sama, contoh: KATAK, APA, KASUR_RUSAK.

#### soal4.go

```go
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	*n = 0
	fmt.Print("Teks : ")
	for *n < NMAX {
		var c rune
		fmt.Scanf(" %c", &c)
		if c == '.' {
			break
		}
		t[*n] = c
		(*n)++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		t[i], t[j] = t[j], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	var temp tabel
	for i := 0; i < n; i++ {
		temp[i] = t[i]
	}
	balikanArray(&temp, n)

	for i := 0; i < n; i++ {
		if t[i] != temp[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	isPalin := palindrom(tab, m)

	balikanArray(&tab, m)

	fmt.Print("Reverse teks : ")
	cetakArray(tab, m)

	fmt.Printf("Palindrom ? %v\n", isPalin)
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul7_9/output/soal4.png)

##### Penjelasan
