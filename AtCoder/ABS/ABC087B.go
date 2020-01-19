package main

import (
	"fmt"
)

func main() {
	// 全探索でないやり方を考える
	var a, b, c, x int
	fmt.Scan(&a, &b, &c, &x)
	count := 0
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			rest := x - i*500 - j*100
			if rest >= 0 && rest%50 == 0 && rest/50 <= c {
				count++
			}
		}
		// var rest,fh,h,f int
		// if x >= 500 {
		// 	rest = x - 500*i
		// } else {
		// 	rest = x
		// }
		// if rest >= 100  {
		// 	rest = rest -
		// }
	}
	fmt.Println(count)
}
