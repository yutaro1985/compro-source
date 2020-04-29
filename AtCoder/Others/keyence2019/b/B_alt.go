package main

import (
	"fmt"
	"strings"
)

// 問題文の条件だと、文字列の最初と最後を見ればいいので
// HasPrefixとHasSuffixを使う
// 参考
// https://atcoder.jp/contests/keyence2019/submissions/4003072

func main() {
	var S string
	fmt.Scan(&S)
	K := "keyence"
	prefixes := make([]string, len(K))
	suffixes := make([]string, len(K))
	for i := 0; i < len(K); i++ {
		prefixes[i] = K[:i]
		suffixes[i] = K[i:]
	}
	for i := 0; i < len(K); i++ {
		if strings.HasPrefix(S, prefixes[i]) && strings.HasSuffix(S, suffixes[i]) {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
