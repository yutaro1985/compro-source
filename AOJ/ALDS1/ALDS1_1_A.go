package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)
	A := make([]int, N)
	for i := range A {
		fmt.Scan(&A[i])
	}
	_ = insSort(A)
}

func insSort(A []int) []int {
	for i := 1; i < len(A); i++ {
		fmt.Println(strings.Trim(fmt.Sprint(A), "[]"))
		v := A[i]
		j := i - 1
		for ; j >= 0 && A[j] > v; j-- {
			A[j+1] = A[j]
		}
		A[j+1] = v
	}
	fmt.Println(strings.Trim(fmt.Sprint(A), "[]"))
	return A
}
