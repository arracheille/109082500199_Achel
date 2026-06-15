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
	fmt.Println("\nBuku ter-favorit:")
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