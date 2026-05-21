package main

import "fmt"

const NMAX int = 1000

type tabInt [NMAX]int
func selectionSort(T *tabInt, n int) {
	var i, j, idxMin, temp int

	for i = 0; i < n-1; i++ {
		idxMin = i

		for j = i + 1; j < n; j++ {
			if T[j] < T[idxMin] {
				idxMin = j
			}
		}

		temp = T[i]
		T[i] = T[idxMin]
		T[idxMin] = temp
	}
}

func main() {
	var n, m int
	var T tabInt

	fmt.Print("Masukan angka :")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {

		fmt.Scan(&m)

		for j := 0; j < m; j++ {
			fmt.Scan(&T[j])
		}

		selectionSort(&T, m)

		for j := 0; j < m; j++ {
			fmt.Print(T[j])

			if j < m-1 {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}