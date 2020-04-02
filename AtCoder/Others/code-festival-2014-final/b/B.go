package main

import "fmt"

func main() {
	var S string
	var ans int
	fmt.Scan(&S)
	// runeをint32に変換するときのやりかた
	// runeはそのままだとUnicodeのCode pointの値なので
	// runeの'0'を引くことによりその差が表している数字と一致するようになる
	for i, s := range S {
		switch i % 2 {
		case 0:
			ans += int(s - '0')
		case 1:
			ans -= int(s - '0')
		}
	}
	fmt.Println(ans)
}
