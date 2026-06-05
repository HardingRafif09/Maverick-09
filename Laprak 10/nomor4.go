package main

import "fmt"

var pita string 
var indeks int  
var currentChar byte

func start() {

	indeks = 0
	currentChar = pita[indeks]
}

func maju() {

	indeks++

	if indeks < len(pita) {
		currentChar = pita[indeks]
	}
}

func eop() bool {

	return currentChar == '.'
}

func cc() byte {

	return currentChar
}

func main() {

	fmt.Print("Masukkan kalimat diakhiri titik : ")
	fmt.Scan(&pita)

	start()

	jumlahKarakter := 0
	jumlahA := 0
	jumlahLE := 0

	var prev byte

	for !eop() {

		fmt.Print(string(cc()))

		jumlahKarakter++

		if cc() == 'A' {
			jumlahA++
		}

		if prev == 'L' && cc() == 'E' {
			jumlahLE++
		}

		prev = cc()

		maju()
	}

	fmt.Println()
	fmt.Println("Jumlah karakter :", jumlahKarakter)

	fmt.Println("Jumlah huruf A :", jumlahA)

	frekuensi := float64(jumlahA) / float64(jumlahKarakter)

	fmt.Println("Frekuensi huruf A :", frekuensi)

	fmt.Println("Jumlah pasangan LE :", jumlahLE)
}