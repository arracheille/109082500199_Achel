# <h1 align="center">Laporan Praktikum Modul 4 - Soal Latihan Modul 4 </h1>

<p align="center">Aqilla Rachel Rabbani - 109082500199</p>

## Unguided

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? (tidak tentunya ya :p)

### Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d.

### Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d.

### Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut!

### P(n, r) = n! / (n−r)!, sedangkan C(n, r) = n! / r!(n−r)!

#### soal1.go

```go
package main

import "fmt"

func factorial (n int, hasil *int) {
	*hasil = 1
	for i := 1; i <= n; i++{
		*hasil *= i
	}
}

func permutation (n,r int, hasil *int) {
	var faktorial_n, faktorial_nr int
	factorial(n, &faktorial_n)
	factorial(n-r, &faktorial_nr)
	*hasil = faktorial_n / faktorial_nr
}

func combination (n,r int, hasil *int) {
	var faktorial_n, faktorial_r, faktorial_nr int
	factorial(n, &faktorial_n)
	factorial(r, &faktorial_r)
	factorial(n-r, &faktorial_nr)
	*hasil = faktorial_n / (faktorial_r * faktorial_nr)
}

func main() {
	var a, b, c, d int
	var p1, c1, p2, c2 int

	fmt.Print("Masukkan nilai variabel a, b, c, d: ")
	fmt.Scan(&a, &b, &c, &d)

	permutation(a, c, &p1)
	combination(a, c, &c1)
	fmt.Println(p1, c1)

	permutation(b, d, &p2)
	combination(b, d, &c2)
	fmt.Print(p2, c2)
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul4/output/output-soal1.png)

##### Penjelasan

Program ini adalah program untuk menghitung faktorial(factorial) dan permutasi(permutation) dari dari variabel a, b, c dan d. Tidak hanya itu, program ini juga menghitung kombinasi(combination) dari faktorial dan permutasi. <strong>Setiap perhitungan faktorial, permutasi, dan kombinasi menggunakan prosedur</strong>. Agar program bekerja dengan baik, nilai variabel a harus lebih besar sama dengan variabel c, dan nilai variabel b lebih besar sama dengan variabel d. Berikut ini penjelasan yang lebih lanjut:

<ul>
  <li><strong>Di dalam func main()</strong>
	<ol>
		<li>Program memiliki variabel a, b, c, dan d dengan tipe data int</li>
		<li>Program memiliki variabel p1, c1, p2, dan c2 dengan tipe data int</li>
		<li>Program menampilkan text perintah untuk memasukkan nilai variabel a, b, c dan d kepada user menggunakan fmt.Print.</li>
		<li>Program memberikan input untuk nilai variabel a, b, c, dan d.Setelah memasukkan input, inputan user dibaca oleh program menggunakan fmt.Scan dan &a, &b, &c, &d, lalu nilai inputan dimasukkan ke variabel a, b, c, dan d secara berurutan.</li>
		<li><strong>Prosedur yang dipanggil & output</strong>. Karena parameter yang digunakan pada prosedur adalah parameter pass by reference, maka setiap parameter yang dipanggil membutuhkan ampersand (&) sebagai parameter aktual.
			<ol>
				<li><strong>Prosedur permutation</strong>, dengan argumen nilai variabel a dan c, lalu &p1 sebagai parameter aktual.</li>
				<li><strong>Prosedur combination</strong>, dengan argumen nilai variabel a dan c, lalu &c1 sebagai parameter aktual.</li>
				<li><strong>Output</strong>, program menggunakan fmt.Println (print di baris baru) untuk menampilkan nilai variabel p1 dan c1</li>
				<li><strong>Prosedur permutation</strong>, dengan argumen nilai variabel b dan d, lalu &p2 sebagai parameter aktual.</li>
				<li><strong>Prosedur combination</strong>, dengan argumen nilai variabel b dan d, lalu &c2 sebagai parameter aktual.</li>
				<li><strong>Output</strong>, program menggunakan fmt.Println (print di baris baru) untuk menampilkan nilai variabel p2 dan c2</li>
			</ol>
		</li>
	</ol>
  </li>
  <li><strong>Prosedur faktorial, permutasi, dan kombinasi</strong>
	<ol>
		<li>Prosedur <strong>factorial</strong>
		</br>Prosedur ini memiliki parameter pass by value yaitu n dengan tipe data integer, dan juga parameter pass by reference yaitu hasil dengan tipe data integer. Karena variabel hasil adalah pointer to int, maka untuk mengaksesnya menggunakan simbol bintang (*). Program memberi nilai 1 ke variabel hasil, lalu membuat perulangan for dengan membuat variabel baru bernama i pada inisiasi dan diberi nilai 1, kondisinya i lebih kecil dari sama dengan variabel n, dan updatenya adalah i++ (post-increment). Kode yang tertulis di dalam for adalah variabel hasil dikalikan dengan variabel i pada setiap perulangan dan nilainya dimasukkan ke variabel hasil. <strong>Singkatnya prosedur ini menghitung faktorial dari suatu nilai menggunakan perulangan for.</strong>
		</li>
		<li>Prosedur <strong>permutation</strong>
		</br>Prosedur ini memiliki parameter pass by value yaitu n dan r dengan tipe data integer, dan juga parameter pass by reference yaitu hasil dengan tipe data integer. Karena variabel hasil adalah pointer to int, maka untuk mengaksesnya menggunakan simbol bintang (*). Prosedur ini membuat variabel baru yaitu faktorial_n dan faktorial_nr dengan tipe data integer. Lalu prosedur memanggil fungsi factorial dengan argumen nilai variabel n dan &faktorial_n sebagai parameter aktual (menyimpan hasil faktorial variabel n), dan memanggil lagi fungsi factorial dengan argumen nilai variabel n dikurangi nilai variabel r dan &faktorial_nr sebagai parameter aktual (menyimpan hasil faktorial variabel n dikurangi variabel r). Terakhir, program menghitung hasil permutasi yaitu nilai variabel faktorial_n dibagi dengan faktorial_nr dan dimasukkan ke variabel hasil.
		</li>
		<li>Prosedur <strong>combination</strong>
		</br>Parameter yang digunakan pada prosedur ini sama dengan prosedur permutation. Prosedur ini membuat variabel baru yaitu faktorial_n, faktorial_r dan faktorial_nr dengan tipe data integer. Lalu prosedur memanggil fungsi factorial dengan argumen nilai variabel n dan &faktorial_n sebagai parameter aktual (menyimpan hasil faktorial variabel n), memanggil fungsi factorial lagi dengan argumen nilai variabel r dan &faktorial_r sebagai parameter aktual (menyimpan hasil faktorial variabel r), dan memanggil fungsi factorial lagi dengan argumen nilai variabel n dikurangi nilai variabel r dan &faktorial_nr sebagai parameter aktual (menyimpan hasil faktorial variabel n dikurangi variabel r). Terakhir, program menghitung hasil kombinasi yaitu nilai variabel faktorial_n dibagi dengan hasil perkalian variabel faktorial_r dengan variabel faktorial_nr dan dimasukkan ke variabel hasil.
		</li>
	</ol>
	<li><strong>Cara kerja Prosedur</strong>
		</br>contohnya permutation(a, c, &p1), maka yang akan terjadi adalah nilai a dimasukkan ke variabel n, nilai c dimasukkan ke variabel r dan akan menjalankan kode prosedur permutation, hasilnya akan dimasukkan ke variabel p1.
	</li>
  </li>
</ul>

### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya.

### Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur.

### prosedure hitungSkor(in/out soal, skor : integer)

### Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit).

### Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.

#### soal2.go

```go
package main

import "fmt"

func hitung_skor(soal *int, skor *int) {
	var menit int
	*soal = 0
	*skor = 0
	for i := 1; i <= 8; i++ {
		fmt.Scan(&menit)
		if menit <= 300 {
			*soal++
			*skor += menit
		}
	}
}

func main() {
	var nama, pemenang string
	var soal, skor, soal_max, skor_menang int

	skor_menang = 999999

	fmt.Println("Masukkan nama, dan 8 nilai menit setiap soal:")
	fmt.Scan(&nama)

	for nama != "Selesai" {
		hitung_skor(&soal, &skor)

		if soal > soal_max || (soal == soal_max && skor < skor_menang) {
			soal_max = soal
			skor_menang = skor
			pemenang = nama
		}

		fmt.Scan(&nama)
	}

	fmt.Println(pemenang, soal_max, skor_menang)
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul4/output/output-soal2.png)

##### Penjelasan

Program ini adalah program untuk menghitung pemenang dari peserta kompetisi pemrograman tingkat nasional dengan syarat pemenang adalah peserta yang menyelesaikan soal paling banyak dengan waktu (menit paling sedikit). Berikut penjelasan yang lebih lanjut:

<ul>
	<li><strong>Di dalam func main()</strong>
		<ol>
			<li>Program memiliki variabel nama dan pemenang dengan tipe data string</li>
			<li>Program memiliki variabel soal, skor, soal_max dan skor_menang dengan tipe data int</li>
			<li>Program memberi nilai ke variabel skor_menang yaitu 999999. Ini adalah nilai skor terbesar.</li>
			<li>Program menampilkan text perintah untuk memasukkan nilai variabel nama menggunakan fmt.Println (print pada baris baru).</li>
			<li>Program memberikan input untuk nilai variabel nama. Setelah memasukkan input, hasil input dibaca oleh program menggunakan fmt.Scan dan &nama, lalu nilai inputan dimasukkan ke variabel nama.</li>
			<li>Program membuat perulangan for dengan kondisi nama tidak sama dengan "Selesai", yang berarti perulangan akan terus berjalan selama user tidak menginputkan teks "Selesai". Isi dari perulangan ini adalah:
				<ol>
					<li>Prosedur hitung_skor yang berisi &soal dan &skor sebagai parameter aktual.</li>
					<li>Perkondisian if, dengan kondisi variabel soal lebih besar daripada variabel soal_max <strong>atau</strong> nilai variabel soal sama dengan soal_max dan variabel skor lebih kecil dari skor_menang. Jika salah satu dari kondisi tersebut terpenuhi, maka nilai variabel soal_max diganti dengan nilai variabel soal, nilai variabel skor_menang diganti dengan nilai variabel skor, dan nilai variabel pemenang diganti dengan nilai variabel nama. <strong>Singkatnya ini adalah kondisi If untuk menemukan siapa pemenang kompetisi</strong></li>
					<li>Terakhir jika user ingin memasukkan data peserta lain, ada input untuk nilai variabel nama. Setelah memasukkan input, hasil input dibaca oleh program menggunakan fmt.Scan dan &nama, lalu nilai inputan dimasukkan ke variabel nama.</li>
				</ol>
			</li>
			<li>Terakhir, program memberikan output pemenang kompetisi yang berisi nama pemenang, soal yang diselesaikan dan skor menang secara berurutan. Output ini menggunakan fmt.Println (print di baris baru)</li>
		</ol>
	</li>
	<li><strong>Prosedur hitung_skor</strong>
	</br>Prosedur ini menggunakan parameter pass by reference yaitu soal dan skor dengan tipe data integer. Karena variabel soal dan skor adalah pointer to int, maka untuk mengaksesnya menggunakan simbol bintang (*). Program membuat variabel baru bernama menit dengan tipe data integer. Lalu program memberi nilai 0 ke masing-masing variabel soal dan skor, lalu membuat perulangan for dengan inisiasi membuat variabel baru bernama i dan diberi nilai 1, kondisinya i lebih kecil dari sama dengan 8, dan updatenya adalah i++ (post-increment). Kode yang tertulis di dalam for adalah input variabel menit, dan sebuah perkondisian if dengan kondisi menit lebih kecil dari 300, jika kondisi ini terpenuhi maka variabel soal akan bertambah 1 pada setiap perulangan, dan variabel skor ditambahkan dengan variabel menit pada setiap perulangan dan nilainya dimasukkan ke variabel skor.</strong>
	</li>
</ul>
