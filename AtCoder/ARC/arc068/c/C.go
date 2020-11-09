package main

import "fmt"

func main() {
	var x, ans int
	fmt.Scan(&x)
	ans = (x / 11) * 2
	if x%11 != 0 {
		if x%11 > 6 {
			ans += 2
		} else {
			ans++
		}
	}
	fmt.Println(ans)
}
