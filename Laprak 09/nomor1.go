package main

import "fmt"

func main() {
	var x, jumlah float64
	var banyak int

	fmt.Scan(&x)

	for x != 9999 {
		jumlah += x
		banyak++
		fmt.Scan(&x)
	}

	if banyak > 0 {
		fmt.Println("Rerata =", jumlah/float64(banyak))
	} else {
		fmt.Println("Tidak ada data")
	}
}