package main

import (
	"fmt"
	"sort"
)

func main() {
	// sortを使ってM番目が1/4M未満かどうかを見る
	var N, M, sum int
	fmt.Scan(&N, &M)
	An := make([]int, N)
	for i := range An {
		fmt.Scan(&An[i])
		sum += An[i]
	}
	sort.Sort(sort.Reverse(sort.IntSlice(An)))
	if An[M-1]*4*M >= sum {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
