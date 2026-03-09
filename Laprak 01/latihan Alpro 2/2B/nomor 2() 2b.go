package main

import "fmt"

func main() {
	var bunga string
	pita := ""
	jumlah := 0

	for i := 1; ; i++ {
		fmt.Print("Bunga ", i, " : ")
		fmt.Scan(&bunga)

		if bunga == "SELESAI" {
			break
		}

		pita = pita + bunga + " - "
		jumlah++
	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", jumlah)
}