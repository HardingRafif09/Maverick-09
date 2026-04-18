package main

import "fmt"

func main() {
	var klubA, klubB string
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	var a, b int
	hasil := []string{}
	i := 1

	for {
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scan(&a, &b)

		if a < 0 || b < 0 {
			break
		}

		if a > b {
			fmt.Printf("Hasil %d : %s\n", i, klubA)
			hasil = append(hasil, klubA)
		} else if b > a {
			fmt.Printf("Hasil %d : %s\n", i, klubB)
			hasil = append(hasil, klubB)
		} else {
			fmt.Printf("Hasil %d : Draw\n", i)
		}
		i++
	}

	fmt.Println("Pertandingan selesai")
}