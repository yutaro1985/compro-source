package main

import (
	"fmt"
	"strings"
)

func main() {
	var N int
	fmt.Scan(&N)
	A := make([]string, N)
	A2 := make([]string, N)
	src := make([]string, N)
	for i := range A {
		fmt.Scan(&A[i])
	}
	copy(A2, A)
	copy(src, A)
	bS := bubbleSort(A, N)
	sS := selectionSort(A2, N)
	fmt.Println(strings.Trim(fmt.Sprint(bS), "[]"))
	fmt.Println(isStable(src, bS))
	fmt.Println(strings.Trim(fmt.Sprint(sS), "[]"))
	fmt.Println(isStable(src, sS))
}

func bubbleSort(A []string, N int) []string {
	for i := 0; i < N; i++ {
		for j := N - 1; j > i; j-- {
			if A[j][1] < A[j-1][1] {
				A[j], A[j-1] = A[j-1], A[j]
			}
		}
	}
	return A
}

func selectionSort(A []string, N int) []string {
	for i := 0; i < N; i++ {
		minj := i
		for j := i; j < N; j++ {
			if A[j][1] < A[minj][1] {
				minj = j
			}
		}
		if A[i][1] != A[minj][1] {
			A[i], A[minj] = A[minj], A[i]
		}
	}
	return A
}

func isStable(in []string, out []string) string {
	N := len(in)
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for a := 0; a < N; a++ {
				for b := a + 1; b < N; b++ {
					if in[i][1] == in[j][1] && in[i] == out[b] && in[j] == out[a] {
						return "Not stable"
					}
				}
			}
		}
	}
	return "Stable"
}
