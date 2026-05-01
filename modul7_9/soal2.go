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