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