package main

import (
	"fmt"
	"sort"
)

// 参考にした回答
// https://atcoder.jp/contests/keyence2020/submissions/9566190
type arms [][]int

func (a arms) Len() int      { return len(a) }
func (a arms) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// 今回はx+lのほうでソートすればよいのでこうしている。
func (a arms) Less(i, j int) bool { return a[i][1] < a[j][1] }

func main() {
	var N int
	fmt.Scan(&N)
	robots := make(arms, N)
	max := -1000000009
	count := 0
	for i := 0; i < N; i++ {
		var x, l int
		fmt.Scan(&x, &l)
		robots[i] = []int{x - l, x + l}
	}
	sort.Sort(robots)
	for i := range robots {
		if robots[i][0] >= max {
			count++
			max = robots[i][1]
		}
	}
	fmt.Println(count)
}
