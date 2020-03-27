package main

import (
	"fmt"
)

// TODO 構造体を使って、複数キーソートでやってみる

func main() {
	var n int
	fmt.Scan(&n)
	cards := make([][]int, 4)
	for i, _ := range cards {
		cards[i] = make([]int, 13)
	}
	for i := 0; i < n; i++ {
		var s string
		var m int
		fmt.Scan(&s, &m)
		m--
		switch s {
		case "S":
			cards[0][m]++
		case "H":
			cards[1][m]++
		case "C":
			cards[2][m]++
		case "D":
			cards[3][m]++
		}
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < 13; j++ {
			if cards[i][j] == 0 {
				switch i {
				case 0:
					fmt.Println("S", j+1)
				case 1:
					fmt.Println("H", j+1)
				case 2:
					fmt.Println("C", j+1)
				case 3:
					fmt.Println("D", j+1)
				}
			}
		}
	}
}
