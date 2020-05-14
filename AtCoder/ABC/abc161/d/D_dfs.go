package main

import (
	"fmt"
	"sort"
)

// https://drken1215.hatenablog.com/entry/2020/04/05/150900
// こちらを参考にした

func main() {
	var K int
	lunlun := make([]int, 0)
	fmt.Scan(&K)
	for i := 1; i <= 9; i++ {
		dfs(1, i, &lunlun)
	}
	sort.Ints(lunlun)
	fmt.Println(lunlun[K-1])
}

func dfs(d, val int, lunlun *[]int) {
	*lunlun = append(*lunlun, val)
	if d == 10 {
		return
	}
	for i := -1; i <= 1; i++ {
		add := val%10 + i
		if 0 <= add && add <= 9 {
			dfs(d+1, val*10+add, lunlun)
		}
	}
}
