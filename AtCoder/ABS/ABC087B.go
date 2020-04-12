package main

import (
	"fmt"
)

func main() {
	// 全探索でないやり方を考える
	var a, b, c, x, rest, ans int
	fmt.Scan(&a, &b, &c, &x)
	for i := 0; i <= a; i++ {
		if x < 500*i {
			break
		} else {
			rest = x - 500*i
		}
		// 100円玉を最も多く使った払い方と、最も少なく使った払い方を求める
		var maxh, minf int
		maxh, minf = max100ycoins(rest, b, c)
		var minh, maxf int
		minh, maxf = min100ycoins(rest, b, c)

		if 500*i+100*maxh+50*minf == x && 500*i+100*minh+50*maxf == x {
			ans += maxh - minh + 1
		}
	}
	fmt.Println(ans)
}

func max100ycoins(rest int, b int, c int) (int, int) {
	var maxh, minf int
	if rest >= 100*b {
		maxh = b
		rest -= 100 * b
	} else {
		maxh = rest / 100
		rest %= 100
	}
	if rest <= c*50 {
		minf = rest / 50
	}
	return maxh, minf
}

func min100ycoins(rest int, b int, c int) (int, int) {
	var minh, maxf int
	if rest >= 50*c {
		if (rest-50*c)%100 != 0 {
			maxf = c - 1
		} else {
			maxf = c
		}
		rest -= 50 * maxf
		minh = rest / 100
	} else {
		maxf = rest / 50
	}
	return minh, maxf
}
