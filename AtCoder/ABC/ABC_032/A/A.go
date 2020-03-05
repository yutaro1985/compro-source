package main

import (
	"fmt"
)

func main() {
	var a, b, n, ans int
	fmt.Scan(&a, &b, &n)
	LCM := lcmof2numbers(a, b)
	if n%LCM != 0 {
		ans = ((n / LCM) + 1) * LCM
	} else {
		ans = n
	}
	fmt.Println(ans)
}

func gcdof2numbers(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcdof2numbers(b, a%b)
}

func lcmof2numbers(a int, b int) int {
	return a * b / gcdof2numbers(a, b)
}
