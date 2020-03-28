package main

import (
	"fmt"
)

// Python3でのやり方を参考にしようと思ったが、Goには値からindexを返す関数がない
// https://onlinejudge.u-aizu.ac.jp/courses/lesson/2/ITP1/6/ITP1_6_B
// 文字をbyteで定義するやり方がcool

func main() {
	var n int
	fmt.Scan(&n)
	soots := []string{"S", "H", "C", "D"}
	cards := make([][]int, 4)
	for i, _ := range cards {
		cards[i] = make([]int, 13)
	}
	for i := 0; i < n; i++ {
		var s string
		var m int
		fmt.Scan(&s, &m)
		m--
		cards[sootCheck(s)][m]++
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			if cards[i][j] == 0 {
				fmt.Println(soots[i], j+1)
			}
		}
	}
}

func sootCheck(s string) int {
	var i int
	switch s {
	case "S":
		i = 0
	case "H":
		i = 1
	case "C":
		i = 2
	case "D":
		i = 3
	}
	return i
}
