package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	fmt.Println(PowMod(m, n, 1e9+7))
}

// PowMod Modular integer power: compute a**b mod m using binary powering algorithm
func PowMod(a, b, m int) int {
	a = a % m
	p := 1 % m
	for b > 0 {
		if b&1 != 0 {
			p = (p * a) % m
		}
		b >>= 1
		a = (a * a) % m
	}
	return p
}
