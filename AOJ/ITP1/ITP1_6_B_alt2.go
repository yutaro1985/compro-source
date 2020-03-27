package main

import "fmt"

// 2次元mapを使った方法
// 他の方の回答を参考に作成

func main() {
	var n int
	fmt.Scan(&n)
	c := make(map[string][]bool)
	suits := []string{"S", "H", "C", "D"}
	for _, s := range suits {
		c[s] = make([]bool, 13)
	}
	for i := 0; i < n; i++ {
		var s string
		var x int
		fmt.Scan(&s, &x)
		c[s][x-1] = true
	}
	for _, s := range suits {
		for i := 0; i < 13; i++ {
			if !c[s][i] {
				fmt.Println(string(s), i+1)
			}
		}
	}
}
