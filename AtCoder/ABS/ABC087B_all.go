package main

import "fmt"

func main() {
	// 愚直な全探索パターン
	var A, B, C, X, ans int
	fmt.Scan(&A, &B, &C, &X)
	for i := 0; i <= A; i++ {
		for j := 0; j <= B; j++ {
			for k := 0; k <= C; k++ {
				if 500*i+100*j+50*k == X {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}
