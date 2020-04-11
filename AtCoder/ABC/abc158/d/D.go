package main

import (
	"fmt"
)

// +=での文字列連結はパフォーマンス的にあまりよろしくないらしい
// https://qiita.com/ruiu/items/2bb83b29baeae2433a79
// https://qiita.com/ono_matope/items/d5e70d8a9ff2b54d5c37
// https://drken1215.hatenablog.com/entry/2020/03/08/144800
// いくつか別解があるらしい
// けんちょんさんの解法2を素直に実装するとTLE
// stringをそのまま扱うとTLEなので、runeにして時間内。ただし、これでも遅い
func main() {
	var S string
	var Q, rev int
	var t, l []rune
	fmt.Scan(&S, &Q)
	Srune := []rune(S)
	for i := 0; i < Q; i++ {
		var T int
		fmt.Scan(&T)
		switch T {
		case 1:
			// 0と1を反転させる
			rev = 1 - rev
		case 2:
			var F int
			var C string
			fmt.Scan(&F, &C)
			Crune := []rune(C)
			F--
			if rev == 1 {
				F = 1 - F
			}
			if F == 0 {
				// https://blog.sarabande.jp/post/89635855383
				t = append(t, Crune...)
			} else {
				l = append(l, Crune...)
			}
		}
	}
	t = Reverserune(t)
	ans := append(t, Srune...)
	ans = append(ans, l...)
	if rev == 1 {
		ans = Reverserune(ans)
	}
	fmt.Println(string(ans))
}

func Reverserune(s []rune) []rune {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
