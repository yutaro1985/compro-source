package main

import "fmt"

func main() {
	var n, m, l int
	fmt.Scan(&n, &m, &l)
	A := make([][]int, n)
	for i := range A {
		A[i] = make([]int, m)
	}
	B := make([][]int, m)
	for i := range B {
		B[i] = make([]int, l)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&A[i][j])
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < l; j++ {
			fmt.Scan(&B[i][j])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < l; j++ {
			var sum int
			for k := 0; k < m; k++ {
				sum += A[i][k] * B[k][j]
			}
			if j != l-1 {
				fmt.Printf("%d ", sum)
			} else {
				fmt.Println(sum)
			}
		}
	}
}
