package main

import (
	"fmt"
)

// ここの解法（いもす法）を参考
// https://drken1215.hatenablog.com/entry/2019/06/11/103300#

func main() {
	var N, M int
	fmt.Scan(&N, &M)
	G := make([]int, N+1)
	ans := 0
	for i := 0; i < M; i++ {
		// IndexとしてはL番目はG[L-1]になるのでマイナス1する
		// いもす法としてはR+1番目を見たいが、IndexとしてはRでよい
		var L, R int
		fmt.Scan(&L, &R)
		L--
		G[L]++
		G[R]--
	}

	for i := 0; i < N; i++ {
		G[i+1] += G[i]
	}

	for i := 0; i < N; i++ {
		if G[i] == M {
			ans++
		}
	}
	fmt.Println(ans)
}
