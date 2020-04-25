package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	var q int
	fmt.Scan(&str, &q)
	s := strings.Split(str, "")
	for i := 0; i < q; i++ {
		var com string
		var a, b int
		fmt.Scan(&com, &a, &b)
		switch com {
		case "print":
			fmt.Println(strings.Join(s[a:b+1], ""))
		case "reverse":
			for j, k := a, b; j < k; j, k = j+1, k-1 {
				s[j], s[k] = s[k], s[j]
			}
		case "replace":
			var rep string
			fmt.Scan(&rep)
			reps := strings.Split(rep, "")
			for j := a; j <= b; j++ {
				s[j] = reps[j-a]
			}
		}
	}
}
