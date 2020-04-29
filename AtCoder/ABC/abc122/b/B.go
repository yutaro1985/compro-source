package main

import "fmt"

func main() {
	var S string
	var ans, cnt int
	fmt.Scan(&S)
	for _, Si := range S {
		if Si == 'A' || Si == 'C' || Si == 'G' || Si == 'T' {
			cnt++
			if cnt > ans {
				ans = cnt
			}
		} else {
			cnt = 0
		}
	}
	fmt.Println(ans)
}
