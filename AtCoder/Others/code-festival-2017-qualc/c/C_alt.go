package main

import (
	"fmt"
)

func main() {
	var s string
	var ans int
	fmt.Scan(&s)
	l, r := 0, len(s)-1
	// 方針が立たなかったので解説を見た
	for l < r {
		if s[l] == s[r] {
			l++
			r--
		} else if s[l] == 'x' {
			l++
			ans++
		} else if s[r] == 'x' {
			r--
			ans++
		} else {
			ans = -1
			break
		}
	}
	fmt.Println(ans)
}
