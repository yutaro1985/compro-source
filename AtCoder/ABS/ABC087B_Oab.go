package main

import "fmt"

func main() {
	var A, B, C, X, ans int
	fmt.Scan(&A, &B, &C, &X)
	for i := 0; i <= A; i++ {
		for j := 0; j <= B; j++ {
			rest := X - 500*i - 100*j
			// 500円と100円を足した額とXの差がプラスであり、
			// 残りの50円玉で作れる（割り切れて、かつ50円玉が足りる
			// ときに条件を満たせる
			if rest >= 0 && rest%50 == 0 && rest/50 <= C {
				ans++
			}
		}
	}
	fmt.Println(ans)
}
