package main

import "fmt"

func main() {
	var c, f int

	fmt.Print("Masukan Celcius : ")
	fmt.Scan(&c)

	fmt.Print("Masukan Fahrenheit : ")
	fmt.Scan(&f)

	r := c * 4 / 5
	k := c + 273


	fmt.Println() 
	fmt.Println("Derajat Celcius:", c)
	fmt.Println("Derajat Reamur:", r)
	fmt.Println("Derajat Fahrenheit:", f)
	fmt.Println("Derajat Kelvin:", k)
}