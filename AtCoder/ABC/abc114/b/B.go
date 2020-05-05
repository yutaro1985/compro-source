package main

import (
	"fmt"
	"strconv"
)

func main() {
	var S string
	fmt.Scan(&S)
	ans := 1000000
	for i := 0; i+2 < len(S); i++ {
		X, _ := strconv.Atoi(S[i : i+3])
		d := 753 - X
		if d < 0 {
			d *= -1
		}
		if ans > d {
			ans = d
		}
	}
	fmt.Println(ans)
}
