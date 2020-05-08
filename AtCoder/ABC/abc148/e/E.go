package main

import (
	"fmt"
)

func main() {
	var N, ans int
	fmt.Scan(&N)
	if N%2 == 0 {
		d := 10
		for d <= N {
			ans += N / d
			d *= 5
		}
	}
	fmt.Println(ans)
}
