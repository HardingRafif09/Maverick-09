package main

import "fmt"

func main() {
	var k int
	var hasil float64

	fmt.Print("Nilai K = ")
	fmt.Scan(&k)

	for i := 0; i <= k; i++ {
		a := float64((4*i + 2) * (4*i + 2))
		b := float64((4*i + 1) * (4*i + 3))

		hasil += a / b
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", hasil)
}