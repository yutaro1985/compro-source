package main

import "fmt"

func main() {
	var N, K, cond int
	fmt.Scan(&N, &K)
	if N%2 != 0 {
		cond = N/2 + 1
	} else {
		cond = N / 2
	}
	if cond >= K {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
