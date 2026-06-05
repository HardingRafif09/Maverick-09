package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Domino struct {
	sisi1 int
	sisi2 int
	nilai int
	balak bool
}

type Dominoes struct {
	kartu  [28]Domino
	jumlah int
}

func buatDomino() Dominoes {

	var d Dominoes
	index := 0

	for i := 0; i <= 6; i++ {

		for j := i; j <= 6; j++ {

			d.kartu[index].sisi1 = i
			d.kartu[index].sisi2 = j

			d.kartu[index].nilai = i + j

			if i == j {
				d.kartu[index].balak = true
			} else {
				d.kartu[index].balak = false
			}

			index++
		}
	}

	d.jumlah = 28

	return d
}

func kocokKartu(d *Dominoes) {

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < d.jumlah; i++ {

		j := rand.Intn(d.jumlah)

		d.kartu[i], d.kartu[j] = d.kartu[j], d.kartu[i]
	}
}

func ambilKartu(d *Dominoes) Domino {

	kartu := d.kartu[d.jumlah-1]

	d.jumlah--

	return kartu
}

func gambarKartu(d Domino, suit int) int {

	if suit == 1 {
		return d.sisi1
	}

	return d.sisi2
}

func nilaiKartu(d Domino) int {
	return d.nilai
}

func main() {

	domino := buatDomino()

	kocokKartu(&domino)

	kartu := ambilKartu(&domino)

	fmt.Println("Kartu yang diambil:")
	fmt.Println(kartu)

	fmt.Println("Gambar sisi 1:", gambarKartu(kartu, 1))

	fmt.Println("Nilai kartu:", nilaiKartu(kartu))
}