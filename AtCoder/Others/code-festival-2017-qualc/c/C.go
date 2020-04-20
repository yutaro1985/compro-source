package main

import (
	"fmt"
)

func main() {
	var s string
	var ans int
	fmt.Scan(&s)
	// 方針が立たなかったので解説を見た
	for s != "" {
		if len(s) == 1 {
			s = ""
		} else if s[0] == s[len(s)-1] {
			s = s[1 : len(s)-1]
		} else {
			if s[0] == 'x' {
				s += "x"
				ans++
			} else if s[len(s)-1] == 'x' {
				s = "x" + s
				ans++
			} else {
				ans = -1
				break
			}
		}
	}
	fmt.Println(ans)
}
