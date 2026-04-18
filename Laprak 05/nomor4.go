package main

import "fmt"

const NMAX = 127

type tabel [NMAX]rune

// input pakai Scanln
func isiArray(t *tabel, n *int) {
	var input string
	fmt.Scanln(&input)

	*n = 0
	for i := 0; i < len(input) && i < NMAX; i++ {
		if input[i] == '.' {
			break
		}
		t[*n] = rune(input[i])
		*n++
	}
}

// cek palindrom
func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks: ")
	isiArray(&tab, &m)

	if palindrom(tab, m) {
		fmt.Println("Palindrom ? true")
	} else {
		fmt.Println("Palindrom ? false")
	}
}