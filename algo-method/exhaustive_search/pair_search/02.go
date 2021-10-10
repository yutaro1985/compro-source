package main

import (
	"fmt"
)

func main() {
	var L, R, ans int
	fmt.Scan(&L, &R)
	for i := L; i <= R; i++ {
		for j := i + 10; j <= R; j += 10 {
			ans++
		}
	}
	fmt.Println(ans)
}
