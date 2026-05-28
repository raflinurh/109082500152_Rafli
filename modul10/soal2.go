package main

import "fmt"

func main() {
	var x, y int
	var ikan [1000]float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	jumlahWadah := x / y
	if x%y != 0 {
		jumlahWadah++
	}

	var totalBeratWadah float64
	var totalSemuaWadah float64

	idxIkan := 0
	for w := 0; w < jumlahWadah; w++ {
		totalBeratWadah = 0
		// Masukkan ikan ke wadah sesuai kapasitas y atau sisa ikan yang ada
		for k := 0; k < y && idxIkan < x; k++ {
			totalBeratWadah += ikan[idxIkan]
			idxIkan++
		}
		totalSemuaWadah += totalBeratWadah
		fmt.Printf("%.2f ", totalBeratWadah)
	}
	fmt.Println()

	rerataWadah := totalSemuaWadah / float64(jumlahWadah)
	fmt.Printf("%.2f\n", rerataWadah)
}