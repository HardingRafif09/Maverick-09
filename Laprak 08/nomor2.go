package main

import "fmt" 

const NMAX int = 1000 

type tabInt [NMAX]int 

func selectionSortAsc(T *tabInt, n int) {
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

func selectionSortDesc(T *tabInt, n int) {
	var i, j, idxMax, temp int

	for i = 0; i < n-1; i++ {
		idxMax = i

		for j = i + 1; j < n; j++ {
			if T[j] > T[idxMax] {
				idxMax = j
			}
		}

		temp = T[i]
		T[i] = T[idxMax]
		T[idxMax] = temp
	}
}

func main() {
	var n, m, x int 
	var ganjil, genap tabInt
	var nganjil, ngenap int

	fmt.Scan(&n) 

	for i := 0; i < n; i++ {

		nganjil = 0 
		ngenap = 0  

		fmt.Scan(&m) 

		for j := 0; j < m; j++ {
			fmt.Scan(&x)

			if x%2 == 1 { 
				ganjil[nganjil] = x
				nganjil++
			} else { 
				genap[ngenap] = x
				ngenap++
			}
		}

		selectionSortAsc(&ganjil, nganjil)
		selectionSortDesc(&genap, ngenap)

		for j := 0; j < nganjil; j++ {
			fmt.Print(ganjil[j], " ")
		}

		for j := 0; j < ngenap; j++ {
			fmt.Print(genap[j])

			if j < ngenap-1 {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}