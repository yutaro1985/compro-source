package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)
	N %= K
	if N < K-N {
		fmt.Println(N)
	} else {
		fmt.Println(K - N)
	}
}
