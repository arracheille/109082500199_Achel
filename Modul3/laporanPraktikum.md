# <h1 align="center">Laporan Praktikum Modul 3 - Soal Latihan Modul 3 </h1>
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

func factorial (n int) int {
	j := 1
	for i := 1; i <= n; i++{
		j *= i
	}
	return j
}

func permutation (n,r int) int {
	p := factorial(n) / factorial(n-r)
	return p
}

func combination (n,r int) int {
	c := factorial(n) / (factorial(r) * factorial(n-r))
	return c
}

func main() {
	var a, b, c, d int

	fmt.Print("Masukkan nilai variabel a, b, c, d: ")
	fmt.Scan(&a, &b, &c, &d)

	fmt.Print(permutation(a, c), " ")
	fmt.Println(combination(a, c))
	
	fmt.Print(permutation(b, d), " ")
	fmt.Print(combination(b, d))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul3/output/output-soal1.png)
##### Penjelasan
Program ini adalah program untuk menghitung faktorial(factorial) dan permutasi(permutation) dari dari variabel a, b, c dan d. Tidak hanya itu, program ini juga menghitung kombinasi(combination) dari faktorial dan permutasi. Setiap perhitungan faktorial, permutasi, dan kombinasi menggunakan function atau fungsi. Agar program bekerja dengan baik, input variabel a harus lebih besar sama dengan variabel c, dan input variabel b lebih besar sama dengan variabel d. Berikut ini penjelasan yang lebih lanjut:
<ul>
  <li><strong>Di dalam func main()</strong>
	<ol>
		<li>Program memiliki variabel a, b c dan d dengan tipe data int</li>
		<li>Program menampilkan text perintah untuk memasukkan nilai variabel a, b, c dan d kepada user menggunakan fmt.Print.</li>
		<li>Program memberikan input untuk nilai variabel a, b, c dan d.Setelah memasukkan input, inputan user dibaca oleh program menggunakan fmt.Scan dan &a, &b, &c, &d, lalu nilai inputan dimasukkan ke variabel a, b, c dan d secara berurutan.</li>
		<li>Program menampilkan output untuk permutasi dari variabel a dan c menggunakan fmt.Print</li>
		<li>Program menampilkan output untuk kombinasi dari variabel a dan c menggunakan fmt.Println, fmt.Println digunakan agar output dibawah output ini berada pada line/baris baru.</li>
		<li>Program menampilkan output untuk permutasi dari variabel b dan d menggunakan fmt.Print</li>
		<li>Program menampilkan output untuk kombinasi dari variabel b dan d menggunakan fmt.Println, fmt.Println digunakan agar output dibawah output ini berada pada line/baris baru.</li>
	</ol>
  </li>
  <li><strong>fungsi/function(func) faktorial, permutasi, dan kombinasi</strong>
	<ol>
		<li><strong>factorial</strong>
		</br>Pada function ini, program menggunakan variabel n, bertipe data integer. Nilai yang dimasukkan ke function dimasukkan ke variabel n. Function ini membuat variabel baru bernama j dan diberi nilai 1, lalu membuat perulangan for dengan membuat variabel baru pada inisiasi, kondisinya i lebih kecil sama dengan n, dan updatenya i++/post-increment. kode yang tertulis di dalam for adalah variabel j dikalikan variabel i pada setiap perulangan dan nilai variabel i dimasukkan ke variabel j. Terakhir, function menggunakan return untuk mengembalikan nilai variabel j. <strong>Singkatnya, function ini menghitung faktorial dari suatu nilai menggunakan perulangan for</strong>
		</li>
		<li><strong>permutation</strong>
		</br>Pada function ini, program menggunakan variabel n dan r, bertipe data integer. Nilai yang ditujukan ke function dimasukkan ke variabel n dan r secara berurutan. Function ini membuat variabel baru bernama p dan didalamnya ada sebuah rumus faktorial dari variabel n dibagi dengan faktorial dari variabel n dikurangi variabel r. Lalu function menggunakan return untuk mengembalikan nilai variabel p.
		</li>
		<li><strong>permutation</strong>
		</br>Sama seperti permutation, program menggunakan variabel n dan r dan nilai dimasukkan ke masing-masing variabel secara berurutan. Function ini membuat variabel baru bernama c dan didalamnya ada sebuah rumus yaitu faktorial dari variabel n dibagi dengan faktorial variabel r, lalu dikali faktorial dari variabel n dikurangi variabel r. Lalu function menggunakan return untuk mengembalikan nilai variabel c.
		</li>
	</ol>
	<li><strong>cara kerja function</strong>
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
![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul3/output/output-soal2.png)
##### Penjelasan
<ul>
  <li>
  <li>
	<ol type="a">
		<li><strong></strong></li>
	</ol>
  </li>
</ul>

### 3. [Lingkaran] Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. 
### Masukan terdiri dari beberapa tiga baris. Baris pertama dan kedua adalah koordinat titik pusat dan radius dari lingkaran 1 dan lingkaran 2, sedangkan baris ketiga adalah koordinat titik sembarang. Asumsi sumbu x dan y dari semua titik dan juga radius direpresentasikan dengan bilangan bulat. 
### Keluaran berupa string yang menyatakan posisi titik "Titik di dalam lingkaran 1 dan 2", "Titik di dalam lingkaran 1", "Titik di dalam lingkaran 2", atau "Titik di luar lingkaran 1 dan 2".

### Fungsi untuk menghitung jarak titik (a, b) dan (c, d) dimana rumus jarak adalah:
### jarak = √(a − c)^2 + (b − d)^2
### dan juga fungsi untuk menentukan posisi sebuah titik sembarang berada di dalam suatu lingkaran atau tidak. function jarak(a,b,c,d : real) -> real {Mengembalikan jarak antara titik (a,b) dan titik (c,d)} function didalam(cx,cy,r,x,y : real) -> boolean {Mengembalikan true apabila titik (x,y) berada di dalam lingkaran yang memiliki titik pusat (cx,cy) dan radius r}
### Catatan: Lihat paket math dalam lampiran untuk menggunakan fungsi math.Sqrt() untuk menghitung akar kuadrat.
#### soal3.go

```go
package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var (
		cx1, cy1, r1 float64
		cx2, cy2, r2 float64
		x, y float64
		output string
	)

	fmt.Print("Masukkan koordinat titik pusat x, y dan radius dari lingkaran 1: ")
	fmt.Scanln(&cx1, &cy1, &r1)
	fmt.Print("Masukkan koordinat titik pusat x, y dan radius dari lingkaran 2: ")
	fmt.Scanln(&cx2, &cy2, &r2)
	fmt.Print("Masukkan koordinat titik sembarang (x, y): ")
	fmt.Scan(&x, &y)

	titik1 := didalam(cx1, cy1, r1, x, y)
	titik2 := didalam(cx2, cy2, r2, x, y)

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
![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul3/output/output-soal3.png)
##### Penjelasan
<ul>
  <li>
  <li>
	<ol type="a">
		<li><strong></strong></li>
	</ol>
  </li>
</ul>