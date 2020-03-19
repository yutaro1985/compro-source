package main

import "fmt"

func main() {
	var N, R, ans int
	fmt.Scan(&N, &R)
	if N < 10 {
		ans = R + 100*(10-N)
	} else {
		ans = R
	}
	fmt.Println(ans)
}
