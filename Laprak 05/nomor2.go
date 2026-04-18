package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Jumlah elemen: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("Isi array:", arr)

	fmt.Print("Indeks ganjil: ")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Print("Indeks genap: ")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var x int
	fmt.Print("Masukkan x: ")
	fmt.Scan(&x)

	fmt.Print("Indeks kelipatan x: ")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var del int
	fmt.Print("Hapus indeks: ")
	fmt.Scan(&del)

	arr = append(arr[:del], arr[del+1:]...)
	fmt.Println("Setelah hapus:", arr)

	sum := 0
	for _, v := range arr {
		sum += v
	}
	mean := float64(sum) / float64(len(arr))
	fmt.Println("Rata-rata:", mean)

	var total float64
	for _, v := range arr {
		total += math.Pow(float64(v)-mean, 2)
	}
	std := math.Sqrt(total / float64(len(arr)))
	fmt.Println("Standar deviasi:", std)

	var cari int
	fmt.Print("Cari frekuensi: ")
	fmt.Scan(&cari)

	count := 0
	for _, v := range arr {
		if v == cari {
			count++
		}
	}
	fmt.Println("Frekuensi:", count)
}