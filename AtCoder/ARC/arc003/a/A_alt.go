package main

import (
	"fmt"
)

func main() {
	var N float64
	var S string
	fmt.Scan(&N, &S)
	// runeを使ってもうちょっとだけスマートに
	scores := make(map[rune]int)
	for _, GPA := range S {
		scores[GPA]++
	}
	ans := float64(scores['A']*4+scores['B']*3+scores['C']*2+scores['D']) / N
	fmt.Println(ans)
}
