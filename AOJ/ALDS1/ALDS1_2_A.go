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
	A, cnt = bubbleSort(A, N)
	fmt.Println(strings.Trim(fmt.Sprint(A), "[]"))
	fmt.Println(cnt)
}

func bubbleSort(A []int, N int) ([]int, int) {
	finished := true
	cnt := 0
	for finished {
		finished = false
		for i := N - 1; i > 0; i-- {
			if A[i] < A[i-1] {
				A[i], A[i-1] = A[i-1], A[i]
				finished = true
				cnt++
			}
		}
	}
	return A, cnt
}
