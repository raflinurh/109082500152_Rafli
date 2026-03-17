package main

import ("fmt"
		"math")

func fx(n int) int {
	return int(math.Pow(float64(n), 2))
}

func gx(n int) int {
	return n - 2
}

func hx(n int) int {
	return n + 1
}

func main() {
	var (
		a, b, c int
	)
	fmt.Scanln(&a, &b, &c)
	fmt.Println(fx(gx(hx(a))))
	fmt.Println(gx(hx(fx(b))))
	fmt.Println(hx(fx(gx(c))))
}
