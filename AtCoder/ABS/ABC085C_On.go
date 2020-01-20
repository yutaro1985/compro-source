// こちらの方法をGoで実装してみる
//https://ikatakos.com/pot/programming_algorithm/contest_history/atcoder/2018/0107_abc085
// O(N)のパターン。鶴亀算の要領
// 1問(b10)だけうまく行かず…。

package main

import "fmt"

func main() {
	var n, y int
	fmt.Scan(&n, &y)
	tmp1000 := 1000 * n
	if tmp1000 > y {
		fmt.Println(-1, -1, -1)
	}
	rest := y - tmp1000
	for i := 0; i <= rest/4000; i++ {
		tmp5000 := i * 4000
		// iが5000円の枚数
		if rest-tmp5000 < 0 {
			break
		}
		if (rest-tmp5000)%9000 == 0 {
			// othは10000円札の枚数。One THousand
			oth := (rest - tmp5000) / 9000
			if oth >= 0 && oth+i <= n {
				fmt.Println(oth, i, n-oth-i)
				return
			}
		}
	}
	fmt.Println(-1, -1, -1)
}
