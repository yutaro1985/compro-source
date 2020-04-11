package main

import (
	"fmt"
	"math"
)

func main() {
	var N, A, B int
	fmt.Scan(&N, &A, &B)
	div := N / (A + B)
	rem := N % (A + B)
	// A==0の場合は結果的に場合分けしなくて良い
	// divが、N個のボールのうち高橋くんが実行する回数（余りを含まない）
	// remがdiv回の操作が行われた余り。
	// remがAより小さかったらそれはすべて青いボール
	// remがAより大きかったらA個のボールが含まれる
	// if A == 0 {
	// 	fmt.Println(0)
	// } else
	fmt.Println(rem, A, MinOf(rem, A), min(rem, A))
	fmt.Println(div*A + MinOf(rem, A))
}

func MinOf(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}

func min(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton min() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Min(float64(res), float64(nums[i])))
	}
	return res
}
