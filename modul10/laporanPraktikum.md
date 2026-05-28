# <h1 align="center">Laporan Praktikum Modul 10 - PENCARIAN NILAI MAX/MIN </h1>
<p align="center">Rafli Nurhidayat - 109082500152</p>


## Unguided 

### 1. Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual. Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual. Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.
#### soal1.go

```go
package main

import "fmt"

func main() {
	var N int
	var berat [1000]float64

	fmt.Scan(&N)

	if N > 1000 {
		N = 1000
	}

	for i := 0; i < N; i++ {
		fmt.Scan(&berat[i])
	}

	if N > 0 {
		min := berat[0]
		max := berat[0]

		for i := 1; i < N; i++ {
			if berat[i] < min {
				min = berat[i]
			}
			if berat[i] > max {
				max = berat[i]
			}
		}
		// Output hasil
		fmt.Printf("%.2f %.2f\n", min, max)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul10/output/output-soal1.png)
Program ini membaca sebuah bilangan bulat N diikuti N bilangan riil yang merupakan berat anak kelinci dan menentukan berat terkecil serta terbesar dari kumpulan data tersebut. Data disimpan dalam sebuah array statis berkapasitas 1000; apabila N > 1000 program membatasi pembacaan hanya pada 1000 elemen pertama. Algoritme pencarian min dan max menginisialisasi min dan max dengan elemen pertama kemudian melakukan satu kali iterasi melalui sisa elemen, memperbarui nilai min atau max ketika ditemui elemen yang lebih kecil atau lebih besar. Setelah pemrosesan, program mencetak dua nilai (min dan max) dengan format dua angka di belakang koma. Kompleksitas waktu adalah O(N) dan penggunaan memori tambahan konstan selain array input.

### 2. Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual. Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual. Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.
#### soal2.go

```go
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
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul10/output/output-soal2.png)
Program ini menerima dua bilangan bulat x (jumlah ikan) dan y (kapasitas tiap wadah), lalu membaca x bilangan riil sebagai berat masing-masing ikan. Program membagi ikan ke dalam wadah-wadah berkapasitas y sesuai urutan input; jika sisa ikan kurang dari y maka wadah terakhir akan berisi sisa tersebut. Untuk setiap wadah dihitung total beratnya dan dicetak berurutan dengan format dua angka desimal. Setelah mencetak total tiap wadah, program menghitung dan menampilkan berat rata-rata per wadah (jumlah semua total wadah dibagi jumlah wadah). Perhitungan jumlah wadah pada kode menggunakan pembagian bulat ditambah satu jika ada sisa (x%y != 0). Algoritme memproses setiap ikan sekali (O(x)) dan membutuhkan memori O(1) tambahan selain array input.

### 3. Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.
#### soal3.go

```go
package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var total float64 = 0
	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}

func main() {
	var n int
	var dataBalita arrBalita
	var min, max float64

	fmt.Print("Masukan banyak data berat balita: ")
	fmt.Scan(&n)

	if n > 100 {
		n = 100
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&dataBalita[i])
	}

	hitungMinMax(dataBalita, n, &min, &max)
	rataRata := rerata(dataBalita, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rataRata)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul10/output/output-soal3.png)
Program ini dirancang untuk pencatatan berat balita (mis. di posyandu). Tipe arrBalita didefinisikan sebagai array statis [100]float64. Fungsi hitungMinMax menerima array dan nilai n bersama dua pointer untuk bMin dan bMax; fungsi ini menginisialisasi kedua nilai dengan elemen pertama lalu melakukan iterasi untuk mencari nilai minimum dan maksimum, memperbarui lokasi memori yang diberikan lewat pointer. Fungsi rerata menjumlahkan semua elemen hingga n dan mengembalikan nilai rata-rata sebagai float64. Pada main program meminta input banyaknya data n (dibatasi maksimum 100), membaca n berat balita, memanggil kedua fungsi tersebut, lalu mencetak hasil minimum, maksimum, dan rata-rata dengan format dua angka desimal dan satuan kg.