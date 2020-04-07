package main

import "fmt"

func main() {
	var S, L, R, ans int
	fmt.Scan(&S, &L, &R)
	if S < L {
		ans = L
	} else if R < S {
		ans = R
	} else {
		ans = S
	}
	fmt.Println(ans)
}
