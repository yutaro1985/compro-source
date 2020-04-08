package main

import (
	"fmt"
)

func main() {
	var A, B, K, cnt int
	fmt.Scan(&A, &B, &K)
	gcd := gcdof2numbers(A, B)
	for i := gcd; ; i-- {
		if gcd%i == 0 {
			cnt++
		}
		if cnt == K {
			fmt.Println(i)
			return
		}
	}
}

func gcdof2numbers(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcdof2numbers(b, a%b)
}
