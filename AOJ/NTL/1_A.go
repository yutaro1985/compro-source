package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	N := n
	ans := make([]int, 0)
	for i := 2; i*i <= N; i++ {
		for N%i == 0 {
			ans = append(ans, i)
			N /= i
		}
	}
	if N != 1 {
		ans = append(ans, N)
	}
	sort.Ints(ans)

	fmt.Println(strconv.Itoa(n)+":", strings.Trim(fmt.Sprint(ans), "[]"))
}
