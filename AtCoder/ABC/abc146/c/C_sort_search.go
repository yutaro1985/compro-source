package main

import (
	"fmt"
	"sort"
	"strconv"
)

// 二分探索の実装例
// sort.Searchを使った例

func main() {
	var A, B, X int
	fmt.Scan(&A, &B, &X)
	//
	N := sort.Search(1e9+1, func(i int) bool {
		// Nが0となることはないのでfalse
		if i == 0 {
			return false
		}
		// 条件を満たすような最小のiを見つけてくれるライブラリなので、
		// 今回求めるのは所持金を上回る購入額となる最小のN。なので答えとしてはN-1を出す
		// 上限＋1まで検索しないと合わなくなる…
		return A*i+B*len(strconv.Itoa(i)) > X
	})
	fmt.Println(N - 1)
}
