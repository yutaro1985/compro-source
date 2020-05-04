package main

import "fmt"

// dfsの練習としてまずは見様見真似実装
// https://drken1215.hatenablog.com/entry/2020/05/04/190252

const M = 2

func main() {
	A := make([]int, 0)
	dfs(A)
}

func dfs(A []int) {
	if len(A) == 10 {
		for _, bi := range A {
			fmt.Printf("%v", bi)
		}
		fmt.Println()
		return
	}

	for i := 0; i < M; i++ {
		A = append(A, i)
		dfs(A)
		A = A[:len(A)-1]
	}
}
