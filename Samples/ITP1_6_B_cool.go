package main

import "fmt"

// 他の方の回答が非常に参考になったのでサンプルとしてコミット
// byteの実態が数字であることを利用して2次元mapを使う発想はお見事

func main() {
	var n int
	fmt.Scan(&n)
	c := make(map[byte][]bool)
	suits := []byte{'S', 'H', 'C', 'D'}
	for _, s := range suits {
		c[s] = make([]bool, 13)
	}
	for i := 0; i < n; i++ {
		var s string
		var x int
		fmt.Scan(&s, &x)
		fmt.Println(s[0])
		c[s[0]][x-1] = true
	}
	for _, s := range suits {
		for i := 0; i < 13; i++ {
			if !c[s][i] {
				fmt.Println(string(s), i+1)
			}
		}
	}
}
