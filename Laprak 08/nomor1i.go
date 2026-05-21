package main

import "fmt"

const NMAX = 1000

type tabInt [NMAX]int

func insertionSort(T *tabInt, n int) {
	var i, pass, temp int

	for pass = 1; pass < n; pass++ {
		temp = T[pass]
		i = pass

		for i > 0 && T[i-1] > temp {
			T[i] = T[i-1]
			i--
		}

		T[i] = temp
	}
}

func cekJarak(T tabInt, n int) {
	var selisih, i int

	if n < 2 {
		fmt.Println("Data berjarak tidak tetap")
		return
	}

	selisih = T[1] - T[0]

	for i = 1; i < n; i++ {
		if T[i]-T[i-1] != selisih {
			fmt.Println("Data berjarak tidak tetap")
			return
		}
	}

	fmt.Println("Data berjarak", selisih)
}

func main() {
	var T tabInt
	var x, n int

	for {
		fmt.Scan(&x)

		if x < 0 {
			break
		}

		T[n] = x
		n++
	}

	insertionSort(&T, n)

	for i := 0; i < n; i++ {
		fmt.Print(T[i])
		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	cekJarak(T, n)
}