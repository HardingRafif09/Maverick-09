package main

import (
	"fmt"
	"math"
)

func main() {
	var pi, suku, lama float64
	i := 0

	for {
		lama = pi

		if i%2 == 0 {
			suku = 1.0 / float64(2*i+1)
		} else {
			suku = -1.0 / float64(2*i+1)
		}

		pi += suku

		if math.Abs(pi-lama) <= 0.00001 {
			break
		}

		i++
	}

	fmt.Printf("Hasil PI: %.10f\n", pi*4)
	fmt.Println("Pada i ke:", i)
}