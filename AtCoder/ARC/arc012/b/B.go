package main

import "fmt"

func main() {
	var N int
	var va, vb, L float64
	fmt.Scan(&N, &va, &vb, &L)
	for i := 0; i < N; i++ {
		L = L / va * vb
	}
	fmt.Printf("%0.7f\n", L)
}
