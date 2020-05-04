package main

import "fmt"

// dfsの練習としてまずは見様見真似実装
// https://drken1215.hatenablog.com/entry/2020/05/04/190252

const (
	M = 3
	N = 5
)

func main() {
	A := make([]int, 0)
	dfs(A)
}

func dfs(A []int) {
	if len(A) == N {
		for _, bi := range A {
			fmt.Printf("%v", bi)
		}
		fmt.Println()
		return
	}

	for i := 0; i < M; i++ {
		// push
		A = append(A, i)
		dfs(A)
		// pop
		A = A[:len(A)-1]
	}
}
