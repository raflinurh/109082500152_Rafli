# <h1 align="center">Laporan Praktikum Modul 12 - SEARCHING </h1>
<p align="center">Rafli Nurhidayat - 109082500152</p>


## Unguided 

### 1. Pada pemilihan ketua RT yang baru saja berlangsung, terdapat 20 calon ketua yang bertanding memperebutkan suara warga. Perhitungan suara dapat segera dilakukan karena warga cukup mengisi formulir dengan nomor dari calon ketua RT yang dipilihnya. Seperti biasa, selalu ada pengisian yang tidak tepat atau dengan nomor pilihan di luar yang tersedia, sehingga data juga harus divalidasi. Tugas Anda untuk membuat program mencari siapa yang memenangkan pemilihan ketua RT. Buatlah program pilkart yang akan membaca, memvalidasi, dan menghitung suara yang diberikan dalam pemilihan ketua RT tersebut. Masukan hanya satu baris data saja, berisi bilangan bulat valid yang kadang tersisipi dengan data tidak valid. Data valid adalah integer dengan nilai di antara 1 s.d. 20 (inklusif). Data berakhir jika ditemukan sebuah bilangan dengan nilai 0. Keluaran dimulai dengan baris berisi jumlah data suara yang terbaca, diikuti baris yang berisi berapa banyak suara yang valid. Kemudian sejumlah baris yang mencetak data para calon apa csaja yang mendapatkan suara.
#### soal1.go

```go
package main

import (
	"fmt"
)

func main() {
	var suara, suaraMasuk, suaraSah int
	var hitungSuara [21]int

	for {
		fmt.Scan(&suara)
		if suara == 0 {
			break
		}
		suaraMasuk++
		if suara >= 1 && suara <= 20 {
			suaraSah++
			hitungSuara[suara]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	for i := 1; i <= 20; i++ {
		if hitungSuara[i] > 0 {
			fmt.Printf("%d:%d\n", i, hitungSuara[i])
		}
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul12/output/output-soal1.png)
Program ini membaca serangkaian bilangan bulat yang diakhiri dengan angka 0, menghitung berapa banyak entri suara yang diterima, memvalidasi suara yang bernilai antara 1 dan 20, serta menghitung banyaknya suara untuk tiap calon. Keluaran mencetak jumlah suara masuk, jumlah suara sah, lalu daftar pasangan nomor calon:jumlah suara untuk setiap calon yang memperoleh suara. Algoritme menghitung frekuensi dengan sebuah array tally ukuran 21 dan berjalan dalam waktu O(m) terhadap panjang input m.

### 2. Berdasarkan program sebelumnya, buat program pilkart yang mencari siapa pemenang pemilihan ketua RT. Sekaligus juga ditentukan bahwa wakil ketua RT adalah calon yang mendapatkan suara terbanyak kedua. Jika beberapa calon mendapatkan suara terbanyak yang sama, ketua terpilih adalah dengan nomor peserta yang paling kecil dan wakilnya dengan nomor peserta terkecil berikutnya. Masukan hanya satu baris data saja, berisi bilangan bulat valid yang kadang tersisipi dengan data tidak valid. Data valid adalah bilangan bulat dengan nilai di antara 1 s.d. 20 (inklusif). Data berakhir jika ditemukan sebuah bilangan dengan nilai 0. Keluaran dimulai dengan baris berisi jumlah data suara yang terbaca, diikuti baris yang berisi berapa banyak suara yang valid. Kemudian tercetak calon nomor berapa saja yang menjadi pasangan ketua RT dan wakil ketua RT yang baru.
#### soal2.go

```go
package main

import (
	"fmt"
)

func main() {
	var suara, suaraMasuk, suaraSah int
	var hitungSuara [21]int

	for {
		fmt.Scan(&suara)
		if suara == 0 {
			break
		}
		suaraMasuk++
		if suara >= 1 && suara <= 20 {
			suaraSah++
			hitungSuara[suara]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)

	ketua := -1
	wakil := -1
	maks1 := -1
	maks2 := -1

	for i := 1; i <= 20; i++ {
		if hitungSuara[i] > maks1 {
			maks2 = maks1
			wakil = ketua
			maks1 = hitungSuara[i]
			ketua = i
		} else if hitungSuara[i] > maks2 {
			maks2 = hitungSuara[i]
			wakil = i
		}
	}

	if ketua != -1 {
		fmt.Printf("Ketua RT: %d\n", ketua)
	}
	if wakil != -1 {
		fmt.Printf("Wakil ketua: %d\n", wakil)
	}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul12/output/output-soal2.png)
Program ini memperluas program penghitungan suara sebelumnya dengan menentukan pasangan ketua dan wakil ketua: calon dengan jumlah suara terbanyak menjadi ketua, dan calon dengan jumlah suara terbanyak kedua menjadi wakil. Jika terdapat nilai sama (seri), aturan tie-breaker memilih nomor calon terkecil sebagai pemenang pertama dan nomor calon terkecil berikutnya sebagai wakil. Program mencetak jumlah suara masuk, jumlah suara sah, lalu nomor ketua dan wakil ketua terpilih. Penentuan pemenang dilakukan dengan satu kali iterasi pada array frekuensi sehingga kompleksitasnya O(1) relatif terhadap jumlah calon tetap (20), atau O(m) termasuk pembacaan input.

### 3. Diberikan n data integer positif dalam keadaan terurut membesar dan sebuah integer lain k, apakah bilangan k tersebut ada dalam daftar bilangan yang diberikan? Jika ya, berikan indeksnya, jika tidak sebutkan "TIDAK ADA". Masukan terdiri dari dua baris. Baris pertama berisi dua buah integer positif, yaitu n dan k. n menyatakan banyaknya data, dimana 1 < n <= 1000000. k adalah bilangan yang ingin dicari. Baris kedua berisi n buah data integer positif yang sudah terurut membesar. Keluaran terdiri dari satu baris saja, yaitu sebuah bilangan yang menyatakan posisi data yang dicari (k) dalam kumpulan data yang diberikan. Posisi data dihitung dimulai dari angka 0. Atau memberikan keluaran "TIDAK ADA" jika data k tersebut tidak ditemukan dalam kumpulan. Program yang dibangun harus menggunakan subprogram dengan mengikuti kerangka yang sudah diberikan berikut ini.
#### soal3.go

```go
package main

import "fmt"

const NMAX = 100000
var data [NMAX]int

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	isiArray(n)

	idx := posisi(n, k)

	if idx != -1 {
		fmt.Println(idx)
	} else {
		fmt.Println("TIDAK ADA")
	}
}

func isiArray(n int) {
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}
}

func posisi(n, k int) int {
	kr := 0
	kn := n - 1

	for kr <= kn {
		med := (kr + kn) / 2
		if data[med] == k {
			return med
		} else if data[med] < k {
			kr = med + 1
		} else {
			kn = med - 1
		}
	}
	return -1
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul12/output/output-soal3.png)
Program ini mengimplementasikan pencarian biner untuk menemukan posisi sebuah bilangan k pada daftar terurut. Program membaca dua bilangan n dan k, kemudian n bilangan terurut, menyimpan data, lalu memanggil fungsi posisi yang mengembalikan indeks (0-based) jika ditemukan atau -1 jika tidak. Keluaran adalah indeks posisi atau teks "TIDAK ADA" bila elemen tidak ditemukan. Kompleksitas pencarian adalah O(log n) dan pembacaan data O(n); ukuran maksimum array dibatasi oleh konstanta NMAX.