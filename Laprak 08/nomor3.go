package main

import "fmt"

const NMAX = 1000000

type tabInt [NMAX]int

func insertionSort(T *tabInt, n int) {
	var i, pass, temp int

	for pass = 1; pass < n; pass++ {
		temp = T[pass]
		i = pass

		for i > 0 && temp < T[i-1] {
			T[i] = T[i-1]
			i--
		}

		T[i] = temp
	}
}

func median(T tabInt, n int) int {
	if n%2 == 1 {
		return T[n/2]
	}

	return (T[(n/2)-1] + T[n/2]) / 2
}

func main() {
	var T tabInt
	var x, n int

	for {
		fmt.Scan(&x)

		if x == -5313 {
			break
		}

		if x == 0 {
			insertionSort(&T, n)
			fmt.Println(median(T, n))
		} else {
			T[n] = x
			n++
		}
	}
}