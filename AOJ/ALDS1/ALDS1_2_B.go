package main

import (
	"fmt"
	"strings"
)

func main() {
	var N, cnt int
	fmt.Scan(&N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}
	A, cnt = selectionSort(A, N)
	fmt.Println(strings.Trim(fmt.Sprint(A), "[]"))
	fmt.Println(cnt)
}

func selectionSort(A []int, N int) ([]int, int) {
	var cnt int
	for i := 0; i < N; i++ {
		minj := i
		for j := i; j < N; j++ {
			if A[j] < A[minj] {
				minj = j
			}
		}
		if A[i] != A[minj] {
			A[i], A[minj] = A[minj], A[i]
			cnt++
		}
	}
	return A, cnt
}
