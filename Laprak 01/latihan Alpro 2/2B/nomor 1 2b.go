package main

import "fmt"

func main() {
	var a1, a2, a3, a4 string
	berhasil := true

	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&a1, &a2, &a3, &a4)
	}

	if a1 != "merah" || a2 != "kuning" || a3 != "hijau" || a4 != "ungu" {
		berhasil = false 
	}

	fmt.Println("Berhasil : ", berhasil)
}