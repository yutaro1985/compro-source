package main

import (
	"fmt"
)

// https://examist.jp/mathematics/integer/totient-function/

func main() {
	var n, num, deno int
	fmt.Scan(&n)
	pf := PrimeFactorize(n)
	num, deno = 1, 1
	for k := range pf {
		num *= k - 1
		deno *= k
	}
	ans := n * num / deno
	fmt.Println(ans)
}

// PrimeFactorize 素因数分解したmapを返す
func PrimeFactorize(n int) map[int]int {
	pf := map[int]int{}
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			pf[i]++
			n /= i
		}
	}
	if n != 1 {
		pf[n]++
	}
	return pf
}
