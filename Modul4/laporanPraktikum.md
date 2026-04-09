# <h1 align="center">Laporan Praktikum Modul 4 - Tugas Modul 4 Prosedur </h1>
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
		<li>Program memiliki variabel a, b, c dan d dengan tipe data int</li>
		<li>Program menampilkan text perintah untuk memasukkan nilai variabel a, b, c dan d kepada user menggunakan fmt.Print.</li>
		<li>Program memberikan input untuk nilai variabel a, b, c dan d.Setelah memasukkan input, inputan user dibaca oleh program menggunakan fmt.Scan dan &a, &b, &c, &d, lalu nilai inputan dimasukkan ke variabel a, b, c dan d secara berurutan.</li>
		<li>Program menampilkan output untuk permutasi dari variabel a dan c menggunakan fmt.Print</li>
		<li>Program menampilkan output untuk kombinasi dari variabel a dan c menggunakan fmt.Println, fmt.Println digunakan agar output selanjutnya berada pada line/baris baru.</li>
		<li>Program menampilkan output untuk permutasi dari variabel b dan d menggunakan fmt.Print</li>
		<li>Program menampilkan output untuk kombinasi dari variabel b dan d menggunakan fmt.Print</li>
	</ol>
  </li>
  <li><strong>Fungsi/function(func) dari faktorial, permutasi, dan kombinasi</strong>
	<ol>
		<li><strong>factorial</strong>
		</br>Pada function ini, program menggunakan variabel n, bertipe data integer. Nilai yang dimasukkan ke function dimasukkan ke variabel n. Function ini membuat variabel baru bernama j dan diberi nilai 1, lalu membuat perulangan for dengan membuat variabel baru pada inisiasi, kondisinya i lebih kecil sama dengan n, dan updatenya i++/post-increment. kode yang tertulis di dalam for adalah variabel j dikalikan variabel i pada setiap perulangan dan nilai variabel i dimasukkan ke variabel j. Terakhir, function menggunakan return untuk mengembalikan nilai variabel j. <strong>Singkatnya, function ini menghitung faktorial dari suatu nilai menggunakan perulangan for.</strong>
		</li>
		<li><strong>permutation</strong>
		</br>Pada function ini, program menggunakan variabel n dan r, bertipe data integer. Nilai yang ditujukan ke function dimasukkan ke variabel n dan r secara berurutan. Function ini membuat variabel baru bernama p dan didalamnya ada sebuah rumus faktorial dari variabel n dibagi dengan faktorial dari variabel n dikurangi variabel r. Lalu function menggunakan return untuk mengembalikan nilai variabel p.
		</li>
		<li><strong>combination</strong>
		</br>Sama seperti permutation, program menggunakan variabel n dan r dan nilai dimasukkan ke masing-masing variabel secara berurutan. Function ini membuat variabel baru bernama c dan didalamnya ada sebuah rumus yaitu faktorial dari variabel n dibagi dengan faktorial variabel r, lalu dikali faktorial dari variabel n dikurangi variabel r. Lalu function menggunakan return untuk mengembalikan nilai variabel c.
		</li>
	</ol>
	<li><strong>Cara kerja function</strong>
		</br>contohnya permutation(a, c), maka yang akan terjadi adalah nilai a dimasukkan ke variabel n, nilai c dimasukkan ke variabel r dan akan menjalankan kode function permutation.
	</li>
  </li>
</ul>

### 2. Diberikan tiga buah fungsi matematika yaitu f (x) = x^2, g (x) = x − 2 dan h (x) = x + 1. Fungsi komposisi (fogoh)(x) artinya adalah f(g(h(x))). Tuliskan f(x), g(x) dan h(x) dalam bentuk function. 
### Masukan terdiri dari sebuah bilangan bulat a, b dan c yang dipisahkan oleh spasi. 
### Keluaran terdiri dari tiga baris. Baris pertama adalah (fogoh)(a), baris kedua (gohof)(b), dan baris ketiga adalah (hofog)(c)!
#### soal2.go

```go
package main

import "fmt"

func f (x int) int {
	fx := x * x
	return fx
}

func g (x int) int {
	gx := x - 2
	return gx
}

func h (x int) int {
	hx := x + 1
	return hx
}

func main() {
	var a, b, c int

	fmt.Print("Masukkan nilai variabel a, b, c: ")
	fmt.Scan(&a, &b, &c)

	fmt.Println(f(g(h(a))))
	fmt.Println(g(h(f(b))))
	fmt.Print(h(f(g(c))))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul4/output/output-soal2.png)
##### Penjelasan
Program ini adalah program untuk menghitung fungsi f, g dan h dari variabel a, b dan c. Setelah ketiga fungsi terhitung, lalu program akan menghitung fungsi fogoh untuk variabel a, gohof untuk variabel b, dan hofog untuk variabel c. Berikut penjelasan yang lebih lanjut:
<ul>
  <li><strong>Di dalam func main()</strong>
	<ol>
		<li>Program memiliki variabel a, b dan c dengan tipe data int</li>
		<li>Program menampilkan text perintah untuk memasukkan nilai variabel a, b dan c kepada user menggunakan fmt.Print.</li>
		<li>Program memberikan input untuk nilai variabel a, b dan c.Setelah memasukkan input, inputan user dibaca oleh program menggunakan fmt.Scan dan &a, &b, &c, lalu nilai inputan dimasukkan ke variabel a, b dan c secara berurutan.</li>
		<li>Program menampilkan output fungsi fogoh dari variabel a menggunakan fmt.Println, fmt.Println digunakan agar output selanjutnya berada pada line/baris baru. f(g(h(a))) berarti program menghitung fungsi h dari nilai variabel a, lalu  menghitung fungsi g menggunakan hasil dari fungsi h, terakhir menghitung fungsi f menggunakan hasil dari fungsi g.</li>
		<li>Program menampilkan output fungsi gohof dari variabel b menggunakan fmt.Println, fmt.Println digunakan agar output selanjutnya berada pada line/baris baru. g(h(f(b))) berarti program menghitung fungsi f dari nilai variabel b, lalu  menghitung fungsi h menggunakan hasil dari fungsi f, terakhir menghitung fungsi g menggunakan hasil dari fungsi h.</li>
		<li>Program menampilkan output fungsi hofog dari variabel c menggunakan fmt.Print. h(f(g(c))) berarti program menghitung fungsi g dari nilai variabel c, lalu  menghitung fungsi f menggunakan hasil dari fungsi g, terakhir menghitung fungsi h menggunakan hasil dari fungsi f.</li>
	</ol>
  </li>
    <li><strong>Fungsi/function(func) dari f, g dan h</strong>
	<ol>
		<li><strong>f</strong>
		</br>Pada function ini, program menggunakan variabel x, bertipe data integer. Nilai yang dimasukkan ke function dimasukkan ke variabel x. Function ini membuat variabel baru bernama fx dan di dalamnya ada sebuah rumus aritmatika yaitu variabel x dikali variabel x. Lalu function menggunakan return untuk mengembalikan nilai variabel fx.
		</li>
		<li><strong>g</strong>
		</br>Pada function ini, program menggunakan variabel x, bertipe data integer. Nilai yang dimasukkan ke function dimasukkan ke variabel x. Function ini membuat variabel baru bernama gx dan di dalamnya ada sebuah rumus aritmatika yaitu variabel x dikurangi 2. Lalu function menggunakan return untuk mengembalikan nilai variabel gx.
		</li>
		<li><strong>h</strong>
		</br>Pada function ini, program menggunakan variabel x, bertipe data integer. Nilai yang dimasukkan ke function dimasukkan ke variabel x. Function ini membuat variabel baru bernama hx dan di dalamnya ada sebuah rumus aritmatika yaitu variabel x ditambah 1. Lalu function menggunakan return untuk mengembalikan nilai variabel hx.
		</li>
	</ol>
	<li><strong>Cara kerja function</strong>
		</br>contohnya f(g(h(a))), maka yang akan terjadi adalah nilai a dimasukkan ke variabel x dari function h, karena perhitungan dimulai dari function yang paling dalam. Lalu setelah program selesai menghitung function h, selanjutnya menghitung fuction g karena function h berada di dalam function g. Perhitungan dari function g adalah hasil dari function h dimasukkan ke variabel x dari function g. Setelah program selesai menghitung, terakhir, karena function g berada di dalam function f, perhitungan dari function f adalah hasil dari function g dimasukkan ke variabel x dari function f.
	</li>
  </li>
</ul>