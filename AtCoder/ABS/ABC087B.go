package main

import (
	"fmt"
)

func main() {
	// 全探索でないやり方を考える
	var a, b, c, x int
	fmt.Scan(&a, &b, &c, &x)
	rest := x
	flag := false
	for i := 0; i <= a; i++ {
		// for j := 0; j <= b; j++ {
		// 	rest := x - i*500 - j*100
		// 	if rest >= 0 && rest%50 == 0 && rest/50 <= c {
		// 		count++
		// 	}
		// }

		// TODO ネストしないで計算する方法を考える
		h := 0
		f := 0
		if rest >= 500 {
			rest = x - 500*i
		} else {
			rest = x
			flag = true
		}
		if rest <= 100*b {
			if rest <= c*50 {
				f += rest / 50
			}
			h += rest / 100
			fmt.Println(h, f)
		} else {
			if rest <= c*50 {
				f += rest / 50
			}
			h += rest / 100
			fmt.Println(h, f)

		}
		if flag {
			fmt.Println(i)
			fmt.Println(i + h + f)
			break
		}
	}
}
