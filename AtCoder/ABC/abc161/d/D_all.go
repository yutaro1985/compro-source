package main

import (
	"fmt"
)

// https://drken1215.hatenablog.com/entry/2020/04/05/150900
// こちらを参考にした

func main() {
	var K int
	lunlun := make([]int, 0)
	cur := make([]int, 0)
	fmt.Scan(&K)
	for i := 1; i <= 9; i++ {
		lunlun = append(lunlun, i)
		cur = append(cur, i)
	}
	for i := 1; i < 10; i++ {
		cur = calc_next(cur)
		lunlun = append(lunlun, cur...)
	}
	fmt.Println(lunlun[K-1])
}

func calc_next(cur []int) []int {
	res := make([]int, 0)
	for _, v := range cur {
		for i := -1; i <= 1; i++ {
			add := v%10 + i
			if add >= 0 && add <= 9 {
				res = append(res, v*10+add)
			}
		}
	}
	return res
}
