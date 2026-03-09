package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&a, &b)

		if a < 0 || b < 0 || a+b > 150 {
			break
		}

		oleng := math.Abs(a-b) >= 9
		fmt.Println("Sepeda motor pak Andi akan oleng:", oleng)
	}

	fmt.Println("Proses selesai.")
}