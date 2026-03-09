package main

import "fmt"

func main() {
	var r float64
	const pi = 3.14

	fmt.Print("Masukan angka : ")
	fmt.Scan(&r)

	volume := (4.0 / 3.0) * pi * r * r * r
	luas := 4.0 * pi * r * r

	fmt.Printf("Bola dengan jari-jari %.0f memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}