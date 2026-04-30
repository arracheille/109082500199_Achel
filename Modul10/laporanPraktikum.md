# <h1 align="center">Laporan Praktikum Modul 10 - Soal Latihan Modul 10 </h1>

<p align="center">Aqilla Rachel Rabbani - 109082500199</p>

## Unguided

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual.

### Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual.

### Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.

#### soal1.go

```go
package main

import "fmt"

type anak_kelinci [1000]float64

func jumlah(jumlah_kelinci *anak_kelinci, n *int) {
	fmt.Print("Masukkan jumlah kelinci: ")
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Printf("Berat kelinci ke-%d (kg): ", i+1)
		fmt.Scan(&jumlah_kelinci[i])
	}
}

func berat(K anak_kelinci, n int) int {
	var index int = 0
	var j int = 1
	for j < n {
		if K[index] < K[j] {
			index = j
		}
		j = j + 1
	}
	return index
}

func ringan(K anak_kelinci, n int) int {
    var index int = 0
    var j int = 1
    for j < n { 
        if K[index] > K[j] {
            index = j
        }
        j = j + 1
    }
    return index
}

func main() {
	var jumlah_kelinci anak_kelinci
	var n int

	jumlah(&jumlah_kelinci, &n)

	terberat := berat(jumlah_kelinci, n)
	terringan := ringan(jumlah_kelinci, n)

	fmt.Println("Kelinci paling berat adalah kelinci dengan berat:", jumlah_kelinci[terberat])
	fmt.Println("Kelinci paling ringan adalah kelinci dengan berat:", jumlah_kelinci[terringan])
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul10/output/output-soal1.png)

##### Penjelasan

Program ini adalah program untuk mencari nilai minimum dan maksimum dari array anak_kelinci. Berikut penjelasan yang lebih lanjut: 

<ul>
    <li>Tipe array <strong>anak_kelinci</strong>
    <br/>Program memiliki array bernama nama_kelinci dengan kapasitas 1000 dan tipe data float64.
    </li>
    <li><strong>Di dalam func main()</strong>
        <ol>
            <li>Program memiliki variabel jumlah_kelinci dengan tipe data anak_kelinci, yang berarti jumlah_kelinci adalah array.</li>
            <li>Program memiliki variabel n dengan tipe data integer</li>
            <li>Program memanggil prosedur jumlah dengan parameter aktual &jumlah kelinci dan &n, & adalah ampersand.</li>
            <li>Program membuat variabel baru bernama terberat menggunakan := dengan value function berat ber-argumen jumlah_kelinci dan n.</li>
            <li>Program membuat variabel baru bernama terringan menggunakan := dengan value function ringan ber-argumen jumlah_kelinci dan n.</li>
            <li>Program memberi output menggunakan fmt.Println (output ditulis di line baru) untuk menampilkan elemen array jumlah_kelinci variabel terberat</li>
            <li>Program memberi output menggunakan fmt.Println (output ditulis di line baru) untuk menampilkan elemen array jumlah_kelinci variabel terringan</li>
        </ol>
    </li>
    <li><strong>Prosedur dan fungsi</strong>
        <ol>
            <li>Prosedur <strong>jumlah</strong>
            </br>Prosedur ini adalah prosedur untuk menampilkan input secara berulang berdasarkan banyaknya jumlah kelinci. Prosedur ini memiliki parameter pass by reference yaitu jumlah_kelinci dengan tipe data anak_kelinci dan n dengan tipe data integer. Kedua parameter diakses menggunakan simbol bintang (*) karena merupakan pointer ke tipe datanya. Di dalam prosedur ini terdapat:
                <ol>
                    <li>Teks menggunakan fmt.Print untuk menampilkan teks perintah memasukkan jumlah kelinci ke user.</li>
                    <li>Input menggunakan fmt.Scan, inputan dimasukkan ke variabel n.</li>
                    <li>Perulangan for dengan inisiasi membuat variabel baru bernama i menggunakan := dan diberi value 0, kondisi variabel i lebih kecil dari variabel n, update i++ (Post-increment).</li>
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

type arrIkan [1000]float64
type arrWadah [1000]float64

func inputIkan(ikan *arrIkan, x *int, y *int) {
	fmt.Print("Masukkan jumlah ikan (x) dan isi per wadah (y): ")
	fmt.Scan(x, y)
	for i := 0; i < *x; i++ {
		fmt.Printf("Berat ikan ke-%d: ", i+1)
		fmt.Scan(&ikan[i])
	}
}

func hitungWadah(ikan arrIkan, x int, y int, wadah *arrWadah) int {
	jumlahWadah := (x + y - 1) / y

	for w := 0; w < jumlahWadah; w++ {
		var total float64 = 0
		for i := w * y; i < (w+1)*y && i < x; i++ {
			total = total + ikan[i]
		}
		wadah[w] = total
	}
	return jumlahWadah
}

func rataWadah(wadah arrWadah, jumlahWadah int) float64 {
	var total float64 = 0
	for i := 0; i < jumlahWadah; i++ {
		total = total + wadah[i]
	}
	return total / float64(jumlahWadah)
}

func beratPerWadah(jumlahWadah *int, wadah *arrWadah){
	fmt.Print("Total berat per wadah: ")
	for i := 0; i < *jumlahWadah; i++ {
		fmt.Printf("%.2f ", wadah[i])
	}
}

func main() {
	var ikan arrIkan
	var wadah arrWadah
	var x, y int

	inputIkan(&ikan, &x, &y)

	jumlahWadah := hitungWadah(ikan, x, y, &wadah)

	beratPerWadah(&jumlahWadah, &wadah)

	fmt.Printf("\nRata-rata berat per wadah: %.2f", rataWadah(wadah, jumlahWadah))
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul10/output/output-soal2.png)

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
                    <li>Perulangan for dengan inisiasi membuat variabel baru bernama i menggunakan := dan diberi value 0, kondisi variabel i lebih kecil dari variabel x, update i++ (Post-increment).</li>
                    <li>Isi dari perulangan ini adalah teks yang menampilkan index berat ikan menggunakan fmt.Printf dan format %d dari variabel i+1. Dan juga input menggunakan fmt.Scan untuk array ikan. Lalu inputan dibaca oleh program menggunakan fmt.Scan dan &ikan[i], lalu nilai inputan dimasukkan ke array ikan berdasarkan index ([i]) secara berurutan.</li>
                </ol>
            </li>
            <li>Fungsi <strong>hitungWadah</strong>
                <br/>Fungsi ini adalah fungsi untuk menghitung total berat ikan di setiap wadah. Fungsi ini memiliki parameter pass by value yaitu ikan dengan tipe data arrIkan, x dan y dengan tipe data integer dan parameter pass by reference yaitu wadah dengan tipe data arrWadah. Parameter wadah diakses menggunakan simbol bintang (*) karena merupakan pointer ke tipe datanya. Tipe data yang dikembalikan adalah integer. Di dalam Fungsi ini terdapat:
                <ol>
                    <li>Variabel total dengan tipe data float64 dan diberi value 0.</li>
                    <li>Membuat variabel baru menggunakan := dengan value (x + y - 1) dibagi y.</li>
                    <li>Perulangan for dengan inisiasi membuat variabel baru bernama w menggunakan := dan diberi value 0, kondisi variabel w lebih kecil dari variabel jumlahWadah, update w++ (Post-increment).</li>
                    <li>Lalu di dalam perulangan tersebut terdapat perulangan for kedua, dengan inisiasi membuat variabel baru bernama i menggunakan := dan diberi value variabel w dikali i, kondisi variabel i lebih kecil dari variabel w ditambah 1 lalu dikali y dan (&&) variabel i lebih kecil dari variabel x, update i++ (Post-increment). Isi dari perulangan ini adalah variabel total dengan value variabel total ditambah elemen array ikan[i]</li>
                    <li>Lalu <strong>dibawah perulangan kedua, dan di dalam perulangan pertama</strong> terdapat array wadah[w] dengan value total</li>
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
                <br/>Prosedur ini adalah prosedur untuk menghitung rata-rata berat per wadah. Prosedur ini memiliki parameter pass by reference yaitu jumlahWadah dengan tipe data integer, dan wadah dengan tipe data arrWadah. Kedua parameter tersebut diakses menggunakan simbol bintang (*) karena merupakan pointer ke tipe datanya. Di dalam prosedur ini terdapat:
                <ol>
                    <li>Teks menggunakan fmt.Print yang bertuliskan "Total berat per wadah: ".</li>
                    <li>Perulangan for dengan inisiasi membuat variabel baru bernama i menggunakan := dan diberi value 0, kondisi variabel i lebih kecil dari variabel jumlahWadah, update i++ (Post-increment).</li>
                    <li>Isi dari perulangan ini adalah teks yang menampilkan total berat per wadah menggunakan fmt.Printf dan format %.2f (float dengan 2 angka di belakang koma) dari array wadah[i].</li>
                </ol>
            </li>
        </ol>
    </li>
</ul>

### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.

### Buatlah program dengan spesifikasi subprogram sebagai berikut:

#### type arrBalita [100]float64
#### func hitungMinMax(arrBerat arrBalita; bMin, bMax *float64) {
#### /* I.S. Terdefinisi array dinamis arrBerat
#### Proses: Menghitung berat minimum dan maksimum dalam array
#### F.S. Menampilkan berat minimum dan maksimum balita */
#### ... }
#### function rerata (arrBerat arrBalita) real {
#### /* menghitung dan mengembalikan rerata berat balita dalam array */
#### ... }

#### soal3.go

```go
package main

import "fmt"

type arrBalita [100]float64

func banyakBalita(arrBerat *arrBalita, n *int) {
	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBerat[i])
	}
}

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	var j int = 1
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]
	for j < n {
		if arrBerat[j] < *bMin {
			*bMin = arrBerat[j]
		}
		if arrBerat[j] > *bMax {
			*bMax = arrBerat[j]
		}
		j = j + 1
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64 = 0
	for i := 0; i < n; i++ {
		total = total + arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var balita arrBalita
	var n int
	var bMin, bMax float64

	banyakBalita(&balita, &n)
	hitungMinMax(balita, n, &bMin, &bMax)

	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg", rerata(balita, n))
}
```

### Output Unguided :

##### Output

![Screenshot Output Unguided 1_1](https://github.com/arracheille/109082500199_Achel/blob/main/Modul10/output/output-soal3.png)

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
                    <li>Perulangan for dengan inisiasi membuat variabel baru bernama i menggunakan := dan diberi value 0, kondisi variabel i lebih kecil dari variabel n, update i++ (Post-increment).</li>
                    <li>Isi dari perulangan ini adalah teks yang menampilkan index berat balita menggunakan fmt.Printf dan format %d (format integer) dari variabel i+1. Dan juga input menggunakan fmt.Scan untuk arrBerat. Lalu tiap hasil input dibaca oleh program menggunakan fmt.Scan dan &arrBerat[i], lalu nilai inputan dimasukkan ke arrBerat berdasarkan index ([i]) secara berurutan.</li>
                </ol>
            </li>
            <li>Prosedur <strong>hitungMinMax</strong>
                <br/>Prosedur ini adalah prosedur untuk menghitung elemen terkecil dan terbesar dari array arrBerat. prosedur ini memiliki parameter pass by value yaitu arrBerat dengan tipe data arrBalita, dan n dengan tipe data integer, serta parameter pass by reference yaitu bMin dan bMax dengan tipe data float64. Parameter pass by reference diakses menggunakan simbol bintang (*) karena merupakan pointer ke tipe datanya. Di dalam prosedur ini terdapat:
                <ol>
                    <li>Variabel j dengan tipe data integer dan diberi value 1.</li>
                    <li>Nilai yang ditunjuk pointer bMin bMin diberi value array ke 0 dari arrBerat.</li>
                    <li>Nilai yang ditunjuk pointer bMin bMax diberi value array ke 0 dari arrBerat.</li>
                    <li>Perulangan for dengan kondisi j lebih kecil dari n</li>
                    <li>Di dalam perulangan tersebut ada if pertama dengan kondisi jika ada elemen di arrBerat[j] yang lebih kecil dari variabel bMin maka value dari variabel bMin adalah elemen arrBerat[j] yang paling kecil tersebut. Ada juga if kedua dengan kondisi jika ada elemen di arrBerat[j] yang lebih besar dari variabel bMax maka value dari variabel bMax adalah elemen arrBerat[j] yang paling besar tersebut. Baris terakhir pada kode for ini adalah value dari variabel j adalah variabel j ditambah 1.</li>
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