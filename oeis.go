package main

import "fmt"

func main() {
	var input, i int
	fmt.Print("Masukkan input: ")
	fmt.Scan(&input)
	fmt.Printf("Hasil dari A000124 Sloane's OEIS dari %d adalah: ", input)
	for i = 0; i < input-1; i++ {
		fmt.Print(oeis(i), "-")
	}
	fmt.Print(oeis(i))
}

func oeis(n int) int {
	return ((n * (n + 1)) / 2) + 1
}
