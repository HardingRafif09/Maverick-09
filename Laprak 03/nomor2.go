package main

import "fmt"

func hitungSkor() (int, int) {
	soal := 0
	skor := 0

	for i := 0; i < 8; i++ {
		var waktu int
		fmt.Scan(&waktu)

		if waktu <= 300 {
			soal++
			skor += waktu
		}
	}

	return soal, skor
}

func main() {
	var nama string
	pemenang := ""
	maxSoal := -1
	minSkor := 999999

	for {
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		soal, skor := hitungSkor()

		if soal > maxSoal || (soal == maxSoal && skor < minSkor) {
			maxSoal = soal
			minSkor = skor
			pemenang = nama
		}
	}

	fmt.Println(pemenang, maxSoal, minSkor)
}