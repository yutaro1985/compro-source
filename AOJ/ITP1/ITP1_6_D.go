package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	Anm := make([][]int, n)
	for i := range Anm {
		Anm[i] = make([]int, m)
	}
	Bn := make([]int, m)
	for i := range Anm {
		for j := range Anm[i] {
			fmt.Scan(&Anm[i][j])
		}
	}
	for i := range Bn {
		fmt.Scan(&Bn[i])
	}
	for i := 0; i < n; i++ {
		var prod int
		for j := 0; j < m; j++ {
			prod += Anm[i][j] * Bn[j]
		}
		fmt.Println(prod)
	}
}
