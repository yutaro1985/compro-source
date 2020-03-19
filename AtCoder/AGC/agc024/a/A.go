package main

import (
	"fmt"
)

func main() {
	var A, B, C, K, ans int
	fmt.Scan(&A, &B, &C, &K)
	if K%2 == 0 {
		ans = A - B
	} else {
		ans = B - A
	}
	if ans > 1e18 || ans < -1e18 {
		fmt.Println("Unfair")
	} else {
		fmt.Println(ans)
	}
}
