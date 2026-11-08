package main

import "fmt"

func gedung() {
	var n int
	fmt.Scan(&n)
	if n <= 40 {
		fmt.Println("TULT")
	} else if 41 <= n && n <= 60 {
		fmt.Println("Tokong Nanas")
	} else if 61 <= n && n <= 100 {
		fmt.Println("Cacuk")
	} else {
		fmt.Println("Kapasitas ruangan tidak mencukupi")
	}
}

func ganjil_genap() {
	var angka int
	var cek_ganjil bool
	fmt.Scan(&angka)
	if 1 <= angka && angka <= 100 {
		cek_ganjil = angka%2 != 0 // bila habis bagi 2 tidak sama 0 maka ganjil
		if cek_ganjil == true {
			fmt.Println("Ganjil")
		} else {
			fmt.Println("Genap")
		}
	}
}

func main() {
	var i int
	for i = 1; i <= 4; i++ {
		ganjil_genap()
	}
}

func piramid() {
	var n, i, j, k int

	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		for j = 1; j <= n-i; j++ {
			fmt.Printf(" ")
		}
		for k = 1; k <= 2*i-1; k++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}
