package main

import "fmt"

func main() {
	var N, K, digits int
	fmt.Scan(&N, &K)
	for N != 0 {
		N /= K
		digits++
	}
	fmt.Println(digits)
}
