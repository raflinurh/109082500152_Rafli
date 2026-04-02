# <h1 align="center">Laporan Praktikum Modul 4 - PROSEDUR </h1>
<p align="center">Rafli Nurhidayat - 109082500152</p>


## Unguided 

### 1. Minggu ini, mahasiswa Fakultas Informatika mendapatkan tugas dari mata kuliah matematika diskrit untuk mempelajari kombinasi dan permutasi. Jonas salah seorang mahasiswa, iseng untuk mengimplementasikannya ke dalam suatu program. Oleh karena itu bersediakah kalian membantu Jonas? Masukan terdiri dari empat buah bilangan asli a, b, c, dan d yang dipisahkan oleh spasi, dengan syarat a ≥ c dan b ≥ d. Keluaran terdiri dari dua baris. Baris pertama adalah hasil permutasi dan kombinasi a terhadap c, sedangkan baris kedua adalah hasil permutasi dan kombinasi b terhadap d. Catatan: permutasi (P) dan kombinasi (C) dari n terhadap r (n ≥ r) dapat dihitung dengan menggunakan persamaan berikut!
#### soal1.go

```go
package main

import "fmt"

func factorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutation(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	p1 := permutation(a, c)
	c1 := combination(a, c)

	p2 := permutation(b, d)
	c2 := combination(b, d)

	fmt.Println(p1, c1)
	fmt.Println(p2, c2)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul4/output/output-soal1.png)
Program ini berfungsi untuk menghitung permutasi dan kombinasi dari dua pasang bilangan. Program menerima empat bilangan asli sebagai input, yaitu a, b, c, dan d dengan syarat a ≥ c dan b ≥ d, kemudian menghitung hasil permutasi dan kombinasi untuk masing-masing pasangan. Terdapat tiga fungsi utama dalam program ini. Pertama, fungsi factorial(n int) yang menghitung nilai faktorial n dengan cara mengalikan semua bilangan dari 1 hingga n secara iteratif. Kedua, fungsi permutation(n, r int) yang menghitung permutasi menggunakan rumus P(n,r) = n! / (n-r)!. Ketiga, fungsi combination(n, r int) yang menghitung kombinasi menggunakan rumus C(n,r) = n! / (r! × (n-r)!). Pada fungsi main, program membaca empat nilai input sekaligus, lalu menghitung permutasi dan kombinasi untuk pasangan pertama (a, c) dan pasangan kedua (b, d). Output ditampilkan dalam dua baris, baris pertama menampilkan permutasi dan kombinasi a terhadap c, sedangkan baris kedua menampilkan permutasi dan kombinasi b terhadap d. Setiap hasil dihitung dengan memanfaatkan fungsi-fungsi yang telah didefinisikan sebelumnya. 

### 2. Kompetisi pemrograman tingkat nasional berlangsung ketat. Setiap peserta diberikan 8 soal yang harus dapat diselesaikan dalam waktu 5 jam saja. Peserta yang berhasil menyelesaikan soal paling banyak dalam waktu paling singkat adalah pemenangnya. Buat program gema yang mencari pemenang dari daftar peserta yang diberikan. Program harus dibuat modular, yaitu dengan membuat prosedur hitungSkor yang mengembalikan total soal dan total skor yang dikerjakan oleh seorang peserta, melalui parameter formal. Pembacaan nama peserta dilakukan di program utama, sedangkan waktu pengerjaan dibaca di dalam prosedur. prosedure hitungSkor(in/out soal, skor : integer) Setiap baris masukan dimulai dengan satu string nama peserta tersebut diikuti dengan adalah 8 integer yang menyatakan berapa lama (dalam menit) peserta tersebut menyelesaikan soal. Jika tidak berhasil atau tidak mengirimkan jawaban maka otomatis dianggap menyelesaikan dalam waktu 5 jam 1 menit (301 menit). Satu baris keluaran berisi nama pemenang, jumlah soal yang diselesaikan, dan nilai yang diperoleh. Nilai adalah total waktu yang dibutuhkan untuk menyelesaikan soal yang berhasil diselesaikan.
#### soal2.go

```go
package main

import (
	"fmt"
)

func hitungSkor(soal *int, skor *int) {
	*soal = 0
	*skor = 0

	for i := 0; i < 8; i++ {
		var waktu int
		fmt.Scan(&waktu)

		if waktu <= 300 {
			*soal++
			*skor += waktu
		}
	}
}

func main() {
	var (
		nama, pemenang     string
		maxSoal, minSkor, totalSoal, totalSkor  int
		pertamaInput = true
	)

	for {
		fmt.Scan(&nama)
		if nama == "Selesai" {
			break
		}

		hitungSkor(&totalSoal, &totalSkor)

		if pertamaInput || totalSoal > maxSoal || (totalSoal == maxSoal && totalSkor < minSkor) {
			pemenang = nama
			maxSoal = totalSoal
			minSkor = totalSkor
			pertamaInput = false
		}
	}

	fmt.Println(pemenang, maxSoal, minSkor)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul4/output/output-soal2.png)
Program ini berfungsi untuk menentukan pemenang kompetisi pemrograman berdasarkan jumlah soal yang berhasil diselesaikan dan total waktu pengerjaannya. Setiap peserta mengerjakan 8 soal, dan hanya soal yang selesai dalam waktu paling lama 300 menit yang dihitung sebagai soal berhasil. Proses perhitungan dibuat secara modular melalui prosedur hitungSkor, yang membaca delapan data waktu pengerjaan peserta lalu mengembalikan jumlah soal yang selesai dan total waktunya melalui parameter keluaran. Pada program utama, nama peserta dibaca terlebih dahulu, kemudian hasil tiap peserta dibandingkan untuk mencari pemenang. Peserta dengan jumlah soal terbanyak akan dipilih sebagai pemenang, sedangkan jika jumlah soal sama maka peserta dengan total waktu paling kecil akan menjadi pemenangnya. Program berhenti ketika nama Selesai dibaca, lalu menampilkan nama pemenang, jumlah soal yang diselesaikan, dan total skor yang diperoleh.

### 3. Skiena dan Revilla dalam Programming Challenges mendefinisikan sebuah deret bilangan. Deret dimulai dengan sebuah bilangan bulat n. Jika bilangan n saat itu genap, maka suku berikutnya adalah 1⁄2n, tetapi jika ganjil maka suku berikutnya bernilai 3n+1. Rumus yang sama digunakan terus menerus untuk mencari suku berikutnya. Deret berakhir ketika suku terakhir bernilai 1. Sebagai contoh jika dimulai dengan n=22, maka deret bilangan yang diperoleh adalah: 22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1 Untuk suku awal sampai dengan 1000000, diketahui deret selalu mencapai suku dengan nilai 1. Buat program skiena yang akan mencetak setiap suku dari deret yang dijelaskan di atas untuk nilai suku awal yang diberikan. Pencetakan deret harus dibuat dalam prosedur cetakDeret yang mempunyai 1 parameter formal, yaitu nilai dari suku awal. prosedure cetakDeret(in n : integer ) Masukan berupa satu bilangan integer positif yang lebih kecil dari 1000000. Keluaran terdiri dari satu baris saja. Setiap suku dari deret tersebut dicetak dalam baris yang dan dipisahkan oleh sebuah spasi.
#### soal3.go

```go
package main

import "fmt"

func cetakDeret(n int) {
	for n != 1 {
		fmt.Print(n, " ")
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Print(1)
}

func main() {
	var n int
	fmt.Scan(&n)
	cetakDeret(n)
	fmt.Println()
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul4/output/output-soal3.png)
Program ini berfungsi untuk menampilkan deret bilangan ala Skiena dan Revilla yang dimulai dari sebuah bilangan awal n hingga mencapai 1. Aturan deretnya adalah jika nilai saat ini genap maka dibagi 2, sedangkan jika ganjil maka dikalikan 3 lalu ditambah 1. Proses tersebut diulang terus-menerus sampai suku terakhir bernilai 1, dan seluruh suku dicetak dalam satu baris dengan pemisah spasi. Untuk menjaga kerapian struktur program, pencetakan deret dipisahkan ke dalam prosedur cetakDeret yang menerima satu parameter berupa nilai awal deret. Pada program utama, nilai awal dibaca dari masukan lalu prosedur dipanggil untuk menampilkan seluruh urutan bilangan sampai selesai.