package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n, ans int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var x int64
		fmt.Scan(&x)
		// こういう意図ではないが…。
		if big.NewInt(x).ProbablyPrime(0) {
			ans++
		}
		// if isPrime(x) {
		// 	ans++
		// }
	}
	fmt.Println(ans)
}

func isPrime(x int) bool {
	if x == 2 {
		return true
	}
	if x <= 1 || x%2 == 0 {
		return false
	}

	for i := 3; i*i <= x; i += 2 {
		if x%i == 0 {
			return false
		}
	}
	return true
}
