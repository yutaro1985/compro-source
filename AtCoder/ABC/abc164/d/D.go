package main

import (
	"fmt"
)

// 解法はこちらを参考
// https://drken1215.hatenablog.com/entry/2020/04/29/171300

func main() {
	var S string
	var ans int
	fmt.Scan(&S)
	val := make([]int, 2019)
	digit := 1
	cur := 0
	val[cur]++
	N := len(S)
	for i := 0; i < N; i++ {
		add := int(S[N-1-i] - '0')
		cur = (cur + digit*add) % 2019
		digit = (digit * 10) % 2019
		val[cur]++
	}
	for _, v := range val {
		ans += v * (v - 1) / 2
	}
	fmt.Println(ans)
}
