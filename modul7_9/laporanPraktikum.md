# <h1 align="center">Laporan Praktikum Modul 7-9 - ARRAY </h1>
<p align="center">Rafli Nurhidayat - 109082500152</p>


## Unguided 

### 1. Suatu lingkaran didefinisikan dengan koordinat titik pusat (cx, cy) dengan radius r. Apabila diberikan dua buah lingkaran, maka tentukan posisi sebuah titik sembarang (x, y) berdasarkan dua lingkaran tersebut. Gunakan tipe bentukan titik untuk menyimpan koordinat, dan tipe bentukan lingkaran untuk menyimpan titik pusat lingkaran dan radiusnya.
#### soal1.go

```go
package main

import (
    "fmt"
    "math"
)

type titik struct {
    x, y float64
}

type lingkaran struct {
    pusat  titik
    radius float64
}

func jarak(p, q titik) float64 {
    return math.Sqrt(math.Pow(p.x-q.x, 2) + math.Pow(p.y-q.y, 2))
}

func didalam(c lingkaran, p titik) bool {
    return jarak(c.pusat, p) <= c.radius
}

func main() {
    var l1, l2 lingkaran
    var p titik

    fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)
    fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)
    fmt.Scan(&p.x, &p.y)

    in1 := didalam(l1, p)
    in2 := didalam(l2, p)

    if in1 && in2 {
        fmt.Println("Titik di dalam lingkaran 1 dan 2")
    } else if in1 {
        fmt.Println("Titik di dalam lingkaran 1")
    } else if in2 {
        fmt.Println("Titik di dalam lingkaran 2")
    } else {
        fmt.Println("Titik di luar lingkaran 1 dan 2")
    }
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul7_9/output/output-soal1.png)
Program ini menentukan posisi sebuah titik terhadap dua lingkaran dengan memanfaatkan tipe bentukan untuk menyimpan koordinat titik dan data lingkaran. Fungsi jarak digunakan untuk menghitung jarak antara dua titik menggunakan rumus Euclidean, sedangkan fungsi didalam mengembalikan true jika jarak titik ke pusat lingkaran lebih kecil atau sama dengan radiusnya. Pada fungsi main, program membaca data kedua lingkaran dan satu titik sembarang, lalu memeriksa apakah titik tersebut berada di dalam lingkaran pertama, lingkaran kedua, keduanya, atau di luar keduanya. Hasil akhirnya ditampilkan sesuai kondisi yang terpenuhi.

### 2. Sebuah array digunakan untuk menampung sekumpulan bilangan bulat. Buatlah program yang digunakan untuk mengisi array tersebut sebanyak N elemen nilai. Asumsikan array memiliki kapasitas penyimpanan data sejumlah elemen tertentu.
#### soal2.go

```go
package main

import (
    "fmt"
    "math"
)

const NMAX int = 100
type tabel [NMAX]int

func main() {
    var data tabel
    var n, x, hapusIdx, cariBil, sum int

    fmt.Print("Masukkan jumlah elemen (N): ")
    fmt.Scan(&n)

    for i := 0; i < n; i++ {
        fmt.Printf("Isi elemen indeks ke-%d: ", i)
        fmt.Scan(&data[i])
    }

    fmt.Print("\na. Isi array: ")
    for i := 0; i < n; i++ {
        fmt.Printf("%d ", data[i])
    }

    fmt.Print("\nb. Indeks ganjil: ")
    for i := 1; i < n; i += 2 {
        fmt.Printf("%d ", data[i])
    }

    fmt.Print("\nc. Indeks genap: ")
    for i := 0; i < n; i += 2 {
        fmt.Printf("%d ", data[i])
    }

    fmt.Print("\nd. Masukkan nilai x untuk kelipatan indeks: ")
    fmt.Scan(&x)
    fmt.Printf("Elemen indeks kelipatan %d: ", x)
    for i := 0; i < n; i++ {
        if i % x == 0 {
            fmt.Printf("%d ", data[i])
        }
    }

    for i := 0; i < n; i++ {
        sum += data[i]
    }
    rata := float64(sum) / float64(n)
    fmt.Printf("\nf. Rata-rata: %.2f", rata)

    var varians float64
    for i := 0; i < n; i++ {
        varians += math.Pow(float64(data[i])-rata, 2)
    }
    sd := math.Sqrt(varians / float64(n))
    fmt.Printf("\ng. Standar Deviasi: %.2f", sd)

    fmt.Print("\nh. Masukkan bilangan yang dicari frekuensinya: ")
    fmt.Scan(&cariBil)
    count := 0
    for i := 0; i < n; i++ {
        if data[i] == cariBil {
            count++
        }
    }
    fmt.Printf("Frekuensi %d: %d kali", cariBil, count)

    fmt.Print("\ne. Masukkan indeks yang ingin dihapus: ")
    fmt.Scan(&hapusIdx)
    if hapusIdx >= 0 && hapusIdx < n {
        for i := hapusIdx; i < n-1; i++ {
            data[i] = data[i+1]
        }
        n-- 
        fmt.Print("Isi array setelah dihapus: ")
        for i := 0; i < n; i++ {
            fmt.Printf("%d ", data[i])
        }
    } else {
        fmt.Println("Indeks tidak valid.")
    }
    fmt.Println()
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 2](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul7_9/output/output-soal2.png)
Program ini mengolah data yang disimpan dalam array bertipe tetap berukuran maksimum 100 elemen. Pengguna terlebih dahulu memasukkan jumlah data dan seluruh isi array, lalu program menampilkan isi array secara lengkap. Setelah itu program memisahkan elemen pada indeks ganjil dan genap, menampilkan elemen yang indeksnya merupakan kelipatan dari nilai x yang dimasukkan pengguna, menghitung rata-rata, menghitung standar deviasi, serta menghitung frekuensi kemunculan sebuah bilangan tertentu. Pada bagian akhir, program juga menyediakan fitur penghapusan elemen berdasarkan indeks yang dipilih dengan cara menggeser elemen-elemen setelahnya ke kiri sehingga data tetap tersusun rapat.

### 3. Sebuah program digunakan untuk menyimpan dan menampilkan nama-nama klub yang memenangkan pertandingan bola pada suatu grup pertandingan. Buatlah program yang digunakan untuk merekap skor pertandingan bola 2 buah klub bola yang berlaga. Pertama-tama program meminta masukan nama-nama klub yang bertanding, kemudian program meminta masukan skor hasil pertandingan kedua klub tersebut. Yang disimpan dalam array adalah nama-nama klub yang menang saja. Proses input skor berhenti ketika skor salah satu atau kedua klub tidak valid (negatif). Di akhir program, tampilkan daftar klub yang memenangkan pertandingan.
#### soal3.go

```go
package main

import "fmt"

func main() {
    var klubA, klubB string
    var skorA, skorB int
    var pemenang [100]string
    var n int = 0

    fmt.Print("Klub A: ")
    fmt.Scan(&klubA)
    fmt.Print("Klub B: ")
    fmt.Scan(&klubB)

    skorA = 0
    skorB = 0

    for skorA >= 0 && skorB >= 0 {
        fmt.Printf("Pertandingan %d: ", n+1)
        fmt.Scan(&skorA, &skorB)

        if skorA >= 0 && skorB >= 0 {
            if skorA > skorB {
                pemenang[n] = klubA
            } else if skorB > skorA {
                pemenang[n] = klubB
            } else {
                pemenang[n] = "Draw"
            }
            n++ 
        }
    }

    fmt.Println()
    for i := 0; i < n; i++ {
        fmt.Printf("Hasil %d: %s\n", i+1, pemenang[i])
    }
    fmt.Println("Pertandingan selesai")
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 3](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul7_9/output/output-soal3.png)
Program ini mencatat hasil pertandingan dua klub secara berulang dan menyimpannya ke dalam array string. Pengguna memasukkan nama klub A dan klub B terlebih dahulu, lalu memasukkan skor kedua klub untuk setiap pertandingan. Selama kedua skor bernilai tidak negatif, program akan menentukan pemenang pertandingan tersebut, yaitu klub A jika skornya lebih besar, klub B jika skornya lebih besar, atau Draw jika skor sama. Hasil tiap pertandingan disimpan di array pemenang, kemudian setelah input negatif diberikan sebagai penanda selesai, program menampilkan seluruh hasil pertandingan yang telah direkam dan menutup dengan pesan bahwa pertandingan selesai.

### 4. Sebuah array digunakan untuk menampung sekumpulan karakter, Anda diminta untuk membuat sebuah subprogram untuk melakukan membalikkan urutan isi array dan memeriksa apakah membentuk palindrom.
#### soal4.go

```go
package main

import "fmt"

const NMAX int = 127
type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
    var k rune
    *n = 0
    for *n < NMAX {
        fmt.Scanf("%c", &k)
        if k == '.' {
            break
        }
        if k != ' ' && k != '\n' && k != '\r' {
            t[*n] = k
            *n++
        }
    }
}

func cetakArray(t tabel, n int) {
    for i := 0; i < n; i++ {
        fmt.Printf("%c ", t[i])
    }
}

func balikanArray(t *tabel, n int) {
    for i := 0; i < n/2; i++ {
        t[i], t[n-1-i] = t[n-1-i], t[i]
    }
}

func palindrom(t tabel, n int) bool {
    asli := t
    balikanArray(&t, n)
    
    for i := 0; i < n; i++ {
        if asli[i] != t[i] {
            return false
        }
    }
    return true
}

func main() {
    var tab tabel
    var m int

    fmt.Print("Teks: ")
    isiArray(&tab, &m)
    isPalin := palindrom(tab, m)
    balikanArray(&tab, m)
    fmt.Print("Reverse teks: ")
    cetakArray(tab, m)

    fmt.Printf("\nPalindrom? %v\n", isPalin)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 4](https://github.com/rflnhdyt626/109082500152_Rafli/blob/main/modul7_9/output/output-soal4.png)
Program ini menyimpan input karakter ke dalam array rune sampai pengguna mengetik tanda titik sebagai penanda akhir input. Fungsi isiArray bertugas membaca karakter satu per satu dan hanya menyimpan karakter yang bukan spasi atau baris baru. Setelah data masuk, fungsi balikanArray menukar elemen dari depan dan belakang agar urutan teks menjadi terbalik, sedangkan fungsi palindrom membandingkan array asli dengan versi terbaliknya untuk menentukan apakah teks tersebut palindrom. Pada fungsi main, program menampilkan hasil reverse teks dan status palindrom dari input yang diberikan.

