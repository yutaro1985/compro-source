package main

import "fmt"

func main() {
	var N, ans int
	fmt.Scan(&N)
	if N%10 < 7 {
		ans = N/10*100 + N%10*15
	} else {
		ans = N/10*100 + 100
	}
	fmt.Println(ans)
}
