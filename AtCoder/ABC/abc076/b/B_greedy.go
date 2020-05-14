package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)
	ans := 1
	for i := 0; i < N; i++ {
		ans = MinOf(ans+K, ans*2)
	}
	fmt.Println(ans)
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
