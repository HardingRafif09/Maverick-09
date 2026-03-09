package main

import "fmt"

func main() {
	var a, b float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&a, &b)

		if a >= 9 || b >= 9 {
			break
		}
	}

	fmt.Println("Proses selesai.")
}