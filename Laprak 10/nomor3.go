package main

import (
	"fmt"
	"math/rand"
	"time"
)
type Domino struct {
	sisi1 int
	sisi2 int
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

func bisaMain(a Domino, b Domino) bool {

	return a.sisi1 == b.sisi1 ||
		a.sisi1 == b.sisi2 ||
		a.sisi2 == b.sisi1 ||
		a.sisi2 == b.sisi2
}

func main() {


	domino := buatDomino()

	kocokKartu(&domino)

	k1 := ambilKartu(&domino)
	k2 := ambilKartu(&domino)

	fmt.Println("Kartu pertama :", k1)
	fmt.Println("Kartu kedua   :", k2)

	if bisaMain(k1, k2) {

		fmt.Println("Kartu bisa disambungkan")

	} else {

		fmt.Println("Kartu tidak bisa disambungkan")
	}
}