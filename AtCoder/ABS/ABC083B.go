package main

import "fmt"

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	//単純計算する以外の方法も考えてみたい
	// oth := 0
	// th := 0
	// h := 0
	// t := 0
	// o := 0
	ans := 0
	for i := 1; i <= n; i++ {
		total := degitsum(i)
		if a <= total && total <= b {
			ans += i
		}
	}
	fmt.Println(ans)
	// if b < 10 {
	// 	// 一桁の数字の数
	// 	total = a - b + 1
	// 	// 二桁以上の数字の数

	// } else {

	// }
}

func degitsum(n int) int {
	var total int
	for {
		total += n % 10
		n /= 10
		if n == 0 {
			break
		}
	}
	return total
}
