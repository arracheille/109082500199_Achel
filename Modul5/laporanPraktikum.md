# <h1 align="center">Laporan Praktikum Modul 5 - Soal Latihan Modul 5 </h1>
<p align="center">Aqilla Rachel Rabbani - 109082500199</p>

## Unguided 

### 1. Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.
#### soal1.go

```go
package main

import "fmt"

func main(){
	var n int

	fmt.Print("Masukkan nilai variabel n: ")
	fmt.Scan(&n)

	fmt.Print("n: ")
	index(0, n)
	fmt.Print("\nSn: ")
	deret(0, n)

	// pada soal tertulis deretnya sampai 10, tapi biar safe saya pakai variabel n, dibawah ini kodenya kalo sampe 10 doang
	// fmt.Print("n: ")
	// index(0, 10)
	// fmt.Print("\nSn: ")
	// deret(0, 10)
}

func index(x, n int){
	if x == n {
		fmt.Print(x)
	} else {
		fmt.Print(x, " ")
		index(x+1, n)
	}
}

func deret(x, n int){
	if x == n {
		fmt.Print(fibonacci(x))
	} else {
		fmt.Print(fibonacci(x), " ")
		deret(x+1, n)
	}
}

func fibonacci(n int) int {
    if n == 0 {
        return 0
    } else if n == 1 {
        return 1
    } else {
        return fibonacci(n-1) + fibonacci(n-2)
    }
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul5/output/output-soal1.png)
##### Penjelasan
Program ini adalah program untuk menghitung deret fibonacci dari dengan deret dari 1 sampai 10. Berikut ini penjelasan yang lebih lanjut:
<ul>
  <li><strong>Di dalam func main()</strong>
	<ol>
		<li>Program memiliki variabel n dengan tipe data integer</li>
		<li>Program menampilkan text perintah untuk memasukkan nilai variabel n ke user menggunakan fmt.Print.</li>
		<li>Program memberikan input untuk nilai variabel n. Setelah memasukkan input, inputan user dibaca oleh program menggunakan fmt.Scan dan &n, lalu nilai inputan dimasukkan ke variabel n.</li>
		<li>Program menampilkan text bertuliskan "n: ", nantinya di sebelahnya akan ada nilai rekursif dari 1 sampai n.</li>
		<li>Program memanggil prosedur index dengan parameter 0 dan variabel n.</li>
		<li>Program menampilkan text bertuliskan "Sn: ", nantinya di sebelahnya akan ada nilai hasil deret fibonacci dari suku ke 1 sampai n.</li>
		<li>Program memanggil prosedur deret dengan parameter 0 dan variabel n.</li>
	</ol>
  </li>
  <li><strong>Prosedur index, deret, dan fungsi fibonacci</strong>. Semua prosedur dan fungsi ini menggunakan rekursif forward.
	<ol>
		<li><strong>Prosedur index</strong>
		</br>Prosedur ini berfungsi untuk menunjukan suku bilangan. Prosedur ini memiliki parameter pass by value yaitu x dan n, bertipe data integer. Prosedur ini berisi base-case, yaitu:
			<ol>
				<li><strong>Kondisi 1 (if)</strong>
				</br>Jika nilai variabel x sama dengan nilai variabel n, maka program akan menjalankan kode untuk menampilkan variabel x menggunakan fmt.Print. Ini adalah kode untuk output terakhir.
				</li>
				<li><strong>Kondisi 2 (else)</strong>
				</br>Jika kondisi pertama tidak terpenuhi atau nilai variabel x tidak sama dengan nilai variabel n, maka program akan menjalankan kode untuk menampilkan variabel x menggunakan fmt.Print, namun di sampingnya ada spasi (" ") karena ini adalah kode untuk menampilkan teks sebelum output terakhir. Lalu kode ini memanggil prosedur index dengan parameter nilai variabel x ditambah 1, dan nilai variabel n, agar perulangan bisa dijalankan.
				</li>
			</ol>
		</li>
		<li><strong>Prosedur deret</strong>
		</br>Prosedur ini berfungsi untuk menunjukan deret fibonacci. Prosedur ini juga memiliki parameter pass by value yaitu x dan n, bertipe data integer. Prosedur ini juga berisi base-case, yaitu:
			<ol>
				<li><strong>Kondisi 1 (if)</strong>
				</br>Jika nilai variabel x sama dengan nilai variabel n, maka program akan menjalankan kode untuk menampilkan output hasil deret fibonacci dari variabel x. Jadi, program menjalankan fungsi fibonacci dari variabel x terlebih dahulu, lalu hasilnya ditampilkan menggunakan fmt.Print. Ini adalah kode untuk menampilkan output terakhir.
				</li>
				<li><strong>Kondisi 2 (else)</strong>
				</br>Jika kondisi pertama tidak terpenuhi atau nilai variabel x tidak sama dengan nilai variabel n, maka program akan menjalankan kode untuk menampilkan variabel hasil deret fibonacci dari variabel x. Jadi, program menjalankan fungsi fibonacci dari variabel x terlebih dahulu, lalu hasilnya ditampilkan menggunakan fmt.Print, namun di sampingnya ada spasi (" ") karena ini adalah kode untuk menampilkan teks sebelum output terakhir. Lalu kode ini memanggil prosedur deret dengan parameter nilai variabel x ditambah 1, dan nilai variabel n, agar perulangan bisa dijalankan.
				</li>
			</ol>
		</li>
		<li><strong>Fungsi fibonacci</strong>
		</br>Prosedur ini berfungsi untuk menghitung bilangan fibonacci. Prosedur ini juga memiliki parameter pass by value yaitu n, bertipe data integer. Prosedur ini juga berisi base-case, yaitu:
			<ol>
				<li><strong>Kondisi 1 (if)</strong>
				</br>Jika nilai variabel n sama dengan 0, maka nilai yang di-return adalah 0.
				</li>
				<li><strong>Kondisi 2 (else-if)</strong>
				</br>Jika nilai variabel n sama dengan 1, maka nilai yang di-return adalah 1.
				</li>
				<li><strong>Kondisi 3 (else)</strong>
				</br>Jika kedua kondisi sebelumnya tidak terpenuhi, maka program akan menjalankan kondisi ini. Yaitu nilai yang di-return adalah hasil perhitungan fungsi fibonacci n-1 ditambahkan hasil perhitungan fungsi fibonacci n-2.
				</li>
			</ol>
		</li>
	</ol>
  </li>
</ul>

### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.
#### soal2.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Baris: ")
	fmt.Scan(&n)
	baris(1, n)
}

func baris(i, n int) {
	if i <= n {
		bintang(i)
		fmt.Println()
		baris(i+1, n)
	}
}

func bintang(jumlah int) {
	if jumlah > 0 {
		fmt.Print("*")
		bintang(jumlah - 1)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul5/output/output-soal2.png)
##### Penjelasan
Program ini adalah program untuk menampilkan banyaknya baris dan bintang berdasarkan nilai n. Berikut penjelasan yang lebih lanjut:
<ul>
  <li><strong>Di dalam func main()</strong>
	<ol>
		<li>Program memiliki variabel n dengan tipe data integer</li>
		<li>Program menampilkan text perintah untuk memasukkan nilai variabel n kepada user menggunakan fmt.Print.</li>
		<li>Program memberikan input untuk nilai variabel n. Setelah memasukkan input, inputan user dibaca oleh program menggunakan fmt.Scan dan &n, lalu nilai inputan dimasukkan ke variabel n.</li>
		<li>Program memanggil prosedur baris dengan argumen 1 dan variabel n.</li>
	</ol>
  </li>
    <li><strong>Prosedur baris dan bintang</strong>. Kedua prosedur ini menggunakan rekursif forward.
	<ol>
		<li><strong>Prosedur baris</strong>
		</br>Prosedur ini berfungsi untuk menampilkan banyaknya baris berdasarkan nilai variabel n. Prosedur ini memiliki parameter pass by value yaitu i dan n, bertipe data integer. Prosedur ini berisi base-case, if dengan kondisi: 
		</br>Jika nilai variabel i kurang dari sama dengan nilai variabel n, maka program akan memanggil fungsi bintang dengan argumen variabel i. Ini berarti program mengirim banyaknya indeks ke prosedur bintang (misalnya indeks ke-1, maka pada prosedur bintang akan menampilkan "*", lalu dilanjutkan ke indeks ke-2, maka pada prosedur bintang akan menampilkan "**", dan seterusnya). Lalu ada fmt.Println() agar setiap baris bintang berada di line output yang baru. Dan program memanggil prosedur baris dengan argumen nilai variabel i ditambah 1. dan variabel n. Kondisi ini akan terus berulang sampai nilai variabel i sama dengan nilai variabel n.
		</li>
		<li><strong>Prosedur bintang</strong>
		</br>Prosedur ini berfungsi untuk menampilkan banyaknya bintang per baris berdasarkan nilai variabel n. Prosedur ini memiliki parameter pass by value yaitu jumlah, bertipe data integer. Prosedur ini berisi base-case, if dengan kondisi: 
		</br>Jika nilai variabel jumlah lebih dari 0, maka program akan menampilkan banyaknya bintang menggunakan fmt.Print("*"). Bintang ini akan terus bertambah ke kanan berdasarkan indeks ke-i. Dan program memanggil prosedur bintang dengan argumen nilai variabel jumlah dikurangi 1.
		</li>
	</ol>
  </li>
</ul>

### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N.
### Masukan terdiri dari sebuah bilangan bulat positif N.
### Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).
#### soal3.go

```go
package main

import "fmt"

func main() {
    var n int
	fmt.Print("Bilangan: ")
    fmt.Scan(&n)
	fmt.Print("Faktor: ")
    faktor(1, n)
}

func faktor(i, n int) {
    if i < n {
        if n % i == 0 {
            fmt.Print(i, ", ")
        }
        faktor(i+1, n)
    } else if i == n {
		fmt.Print(i)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul5/output/output-soal3.png)
##### Penjelasan
Program ini adalah program untuk menghitung/mencari faktor dari n. Berikut penjelasan yang lebih lanjut:
<ul>
  <li><strong>Di dalam func main()</strong>
	<ol>
		<li>Program memiliki variabel n dengan tipe data integer</li>
		<li>Program menampilkan text perintah untuk memasukkan nilai variabel n kepada user menggunakan fmt.Print.</li>
		<li>Program memberikan input untuk nilai variabel n. Setelah memasukkan input, inputan user dibaca oleh program menggunakan fmt.Scan dan &n, lalu nilai inputan dimasukkan ke variabel n.</li>
		<li>Program menampilkan text bertuliskan "Faktor: ", nantinya di sebelahnya akan ada nilai hasil pencarian faktor dari variabel n.</li>
		<li>Program memanggil prosedur faktor dengan argumen 1 dan variabel n.</li>
	</ol>
	</li>
		<li><strong>Prosedur faktor</strong>. Prosedur ini menggunakan rekursif forward.
		</br>Prosedur ini berfungsi untuk mencari apa saja faktor dari nilai variabel n. Prosedur ini memiliki parameter pass by value yaitu i dan n, bertipe data integer. Prosedur ini berisi base-case dengan kondisi: 
			<ol>
				<li><strong>Kondisi 1 (if)</strong>
				</br>Jika nilai variabel i kurang dari nilai variabel n, maka program akan menjalankan sebuah kondisi untuk mengecek apakah n adalah faktor dari i. Kondisinya adalah jika nilai variabel n habis dibagi atau di-modulus nilai variabel i hasilnya 0, maka program akan menampilkan nilai variabel i tersebut dan titik spasi di sebelahnya, karena ini adalah kode sebelum nilai faktor terakhir
				</li>
				<li><strong>Kondisi 2 (else-if)</strong>
				</br>Jika nilai variabel i sama dengan n, maka program akan menampilkan nilai i tersebut. Kode ini tidak ada titik spasi karena ini adalah faktor terakhir.
				</li>
			</ol>
	</li>
</ul>
