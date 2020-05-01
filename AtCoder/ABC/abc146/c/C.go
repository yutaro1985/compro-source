package main

import (
	"fmt"
	"strconv"
)

// 二分探索の実装例
// めぐる式二分探索
// https://twitter.com/meguru_comp/status/697008509376835584
// https://qiita.com/drken/items/97e37dd6143e33a64c8c

func main() {
	var A, B, X int
	start := 0
	end := int(1e9)
	fmt.Scan(&A, &B, &X)
	for end-start > 1 {
		i := (end + start) / 2
		digits := len(strconv.Itoa(i))
		N := i*A + digits*B
		if N <= X {
			start = i
		} else {
			end = i
		}
	}
	if A*end+B*len(strconv.Itoa(end)) <= X {
		fmt.Println(end)
	} else {
		fmt.Println(start)
	}
}
