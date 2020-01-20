package main

import "fmt"

func main() {
	var n, y int
	fmt.Scan(&n, &y)
	for i := 0; i <= n; i++ {
		for j := 0; i+j <= n; j++ {
			if 10000*i+5000*j+1000*(n-i-j) == y {
				fmt.Println(i, j, n-i-j)
				return
			}
		}
	}
	fmt.Println(-1, -1, -1)
}
