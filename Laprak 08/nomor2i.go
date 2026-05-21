package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating      int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(p *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&p[i].id, &p[i].judul, &p[i].penulis, &p[i].penerbit,
			&p[i].eksemplar, &p[i].tahun, &p[i].rating)
	}
}

func CetakTerfavorit(p DaftarBuku, n int) {
	idx := 0

	for i := 1; i < n; i++ {
		if p[i].rating > p[idx].rating {
			idx = i
		}
	}

	fmt.Println(p[idx].judul, p[idx].penulis, p[idx].penerbit, p[idx].tahun)
}

func UrutBuku(p *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		temp := p[i]
		j := i

		for j > 0 && p[j-1].rating < temp.rating {
			p[j] = p[j-1]
			j--
		}

		p[j] = temp
	}
}

func Cetak5Terbaru(p DaftarBuku, n int) {
	limit := 5
	if n < 5 {
		limit = n
	}

	for i := 0; i < limit; i++ {
		fmt.Println(p[i].judul)
	}
}

func CariBuku(p DaftarBuku, n, r int) {
	low, high := 0, n-1

	for low <= high {
		mid := (low + high) / 2

		if p[mid].rating == r {
			fmt.Println(p[mid].judul, p[mid].penulis, p[mid].penerbit,
				p[mid].tahun, p[mid].eksemplar, p[mid].rating)
			return
		} else if p[mid].rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	fmt.Println("Tidak ada buku dengan rating seperti itu")
}

func main() {
	var p DaftarBuku
	var n, r int

	fmt.Scan(&n)

	DaftarkanBuku(&p, n)

	CetakTerfavorit(p, n)

	UrutBuku(&p, n)

	Cetak5Terbaru(p, n)

	fmt.Scan(&r)
	CariBuku(p, n, r)

}