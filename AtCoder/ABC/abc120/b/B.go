package main

import (
	"fmt"
)

func main() {
	var A, B, K, cnt, ans int
	fmt.Scan(&A, &B, &K)
	gcd := gcdof2numbers(A, B)
	for i := gcd; ; i-- {
		if A%i == 0 && B%i == 0 {
			cnt++
		}
		if cnt == K {
			ans = i
			break
		}
	}
	fmt.Println(ans)
}

func gcdof2numbers(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcdof2numbers(b, a%b)
}
