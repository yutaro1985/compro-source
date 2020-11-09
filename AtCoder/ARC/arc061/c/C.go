package main

import (
	"fmt"
	"strconv"
)

// dfsの実装を見様見真似でやってみる

func main() {
	var S string
	fmt.Scan(&S)
	fmt.Println(dfs(S, 0, uint(len(S)-1)))
}

func dfs(S string, f int, index uint) int {
	if int(index) == len(S)-1 {
		v, j := 0, 0
		for i := range S {
			if f&1 == 1 {
				n, _ := strconv.Atoi(S[j : i+1])
				v += n
				j = i + 1
			}
			f >>= 1
		}
		n, _ := strconv.Atoi(S[j:])
		v += n
		return v
	}
	return dfs(S, f, index+1) + dfs(S, f|(1<<(index+1)), index+1)
}
