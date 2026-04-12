# <h1 align="center">Laporan Praktikum Modul 5 - REKURSIF </h1>
<p align="center">Rafli Nurhidayat - 109082500152</p>


## Unguided 

### 1. Deret fibonacci adalah sebuah deret dengan nilai suku ke-0 dan ke-1 adalah 0 dan 1, dan nilai suku ke-n selanjutnya adalah hasil penjumlahan dua suku sebelumnya. Secara umum dapat diformulasikan Sn = Sn−1 + Sn−2 . Berikut ini adalah contoh nilai deret fibonacci hingga suku ke-10. Buatlah program yang mengimplementasikan fungsi rekursif pada deret fibonacci tersebut.
#### soal1.go

```go
package main

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func main() {
	var n int

	fmt.Scanln(&n)
	fmt.Println(fib(n)) 
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul5/output/output-soal1.png)
Program ini berfungsi untuk menghitung bilangan Fibonacci ke-n menggunakan pemanggilan fungsi rekursif, di mana pengguna memasukkan sebuah nilai n dan program akan menampilkan nilai Fibonacci yang sesuai.

### 2. Buatlah sebuah program yang digunakan untuk menampilkan pola bintang berikut ini dengan menggunakan fungsi rekursif. N adalah masukan dari user.
#### soal2.go

```go
package main

import "fmt"

func printBintang(n int) {
	if n <= 0 {
		return
	}
	printBintang(n - 1)

	for i := 0; i < n; i++ {
		fmt.Print("*")
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Scanln(&n)
	printBintang(n)
}


```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul5/output/output-soal2.png)
Program ini berfungsi untuk menampilkan pola bintang berbentuk segitiga secara bertahap dari 1 hingga N baris menggunakan fungsi rekursif. Pengguna memasukkan sebuah bilangan N, kemudian program akan mencetak baris pertama dengan 1 bintang, baris kedua dengan 2 bintang, dan seterusnya sampai baris ke-N dengan jumlah bintang sebanyak N.

### 3. Buatlah program yang mengimplementasikan rekursif untuk menampilkan faktor bilangan dari suatu N, atau bilangan yang apa saja yang habis membagi N. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan yang menjadi faktor dari N (terurut dari 1 hingga N ya).
#### soal3.go

```go
package main

import "fmt"

func printFaktor(n int, i int) {
	if i > n {
		return
	}
	if n%i == 0 {
		fmt.Print(i, " ")
	}
	printFaktor(n, i+1)
}

func main() {
	var n int
	fmt.Scanln(&n)
	printFaktor(n, 1)
	fmt.Println()
}

```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul5/output/output-soal3.png)
Program ini berfungsi untuk menampilkan semua faktor dari sebuah bilangan bulat positif N menggunakan rekursi, yaitu dengan memeriksa setiap bilangan dari 1 hingga N dan mencetak bilangan yang habis membagi N.

### 4. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan tertentu. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan dari N hingga 1 dan kembali ke N.
#### soal4.go

```go
package main

import "fmt"

func printBarisan(n int) {
	if n == 0 {
		return
	}
	fmt.Print(n, " ")
	printBarisan(n - 1)
	fmt.Print(n, " ")
}

func main() {
	var n int
	fmt.Scanln(&n)
	printBarisan(n)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul5/output/output-soal4.png)
Program ini berfungsi untuk menampilkan barisan bilangan secara rekursif dengan pola menurun dari n hingga 1, lalu dilanjutkan kembali secara menaik dari 1 hingga n, sehingga menghasilkan urutan yang simetris.

### 5. Buatlah program yang mengimplementasikan rekursif untuk menampilkan barisan bilangan ganjil. Masukan terdiri dari sebuah bilangan bulat positif N. Keluaran terdiri dari barisan bilangan ganjil dari 1 hingga N.
#### soal5.go

```go
package main

import "fmt"
func printGanjil(n int) {
	if n <= 0 {
		return
	}
	printGanjil(n - 1)
	if n%2 != 0 {
		fmt.Print(n, " ")
	}
}

func main() {
	var n int
	fmt.Scanln(&n)
	printGanjil(n)
	fmt.Println()
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul5/output/output-soal5.png)
Program ini berfungsi untuk menampilkan semua bilangan ganjil dari 1 hingga n menggunakan rekursi. Program menerima sebuah bilangan bulat sebagai input, kemudian fungsi printGanjil akan memanggil dirinya sendiri dari nilai terkecil sampai n dan mencetak hanya bilangan yang bernilai ganjil secara berurutan.

### 6. Buatlah program yang mengimplementasikan rekursif untuk mencari hasil pangkat dari dua buah bilangan. Masukan terdiri dari bilangan bulat x dan y. Keluaran terdiri dari hasil x dipangkatkan y. Catatan: diperbolehkan menggunakan asterik "*", tapi dilarang menggunakan import "math".
#### soal6.go

```go
package main

import "fmt"

func pangkat(x, y int) int {
	if y == 0 {
		return 1
	}
	return x * pangkat(x, y-1)
}

func main() {
	var x, y int
	fmt.Scanln(&x, &y)
	fmt.Println(pangkat(x, y))
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul5/output/output-soal6.png)
Program ini berfungsi untuk menghitung hasil pemangkatan dua bilangan bulat, yaitu x dipangkatkan y, dengan menggunakan fungsi rekursif. Program menerima input nilai x dan y, lalu memprosesnya dengan kasus dasar saat y bernilai 0 sehingga hasilnya 1, sedangkan untuk nilai y lainnya program akan mengalikan x secara berulang melalui pemanggilan fungsi terhadap y-1 hingga diperoleh hasil akhir, kemudian menampilkannya ke layar.

