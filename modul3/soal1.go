package main

import ("fmt"
)

func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * faktorial(n-1)
}

func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}


func main() {
	var (
		a, b, c, d int
	)
	fmt.Scanln(&a, &b, &c, &d)
	fmt.Print(permutasi(a, c), " ")
	fmt.Println(kombinasi(a, c))

	fmt.Print(permutasi(b, d), " ")
	fmt.Println(kombinasi(b, d))
}