package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	vars := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&vars[i])
	}
	fmt.Println(minOf(vars), maxOf(vars), sum(vars))
}

func maxOf(vars []int) int {
	sort.Ints(vars)
	return vars[len(vars)-1]
}

func minOf(vars []int) int {
	sort.Ints(vars)
	return vars[0]
}

func sum(vars []int) int {
	var sum int
	for _, num := range vars {
		sum += num
	}
	return sum
}
