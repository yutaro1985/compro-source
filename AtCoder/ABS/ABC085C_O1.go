// TODO #9 こちらの方法をGoで実装してみる
//https://ikatakos.com/pot/programming_algorithm/contest_history/atcoder/2018/0107_abc085
// O(1)のパターン

package main

import (
	"fmt"
	"math"
)

func main() {
	var N, Y, x, y, z int
	fmt.Scan(&N, &Y)
	Y /= 1000
	// fmt.Println(N, Y, int(math.Ceil(float64(5*N-Y))/4), int(math.Trunc(float64(10*N-Y))/9))
	if int(math.Ceil(float64(5*N-Y))/4) > MinOf(N, Y) || int(math.Trunc(float64(10*N-Y))/9) < 0 {
		fmt.Println(-1, -1, -1)
		return
	}
	if int(math.Ceil(math.Ceil(float64(5*N-Y))/4)/4) >= 0 {
		if int(math.Ceil(math.Ceil(float64(5*N-Y))/4)/4)%5 == Y%5 {
			z = int(math.Ceil(math.Ceil(float64(5*N-Y)) / 4))
		} else {
			z = Y % 5 * (int(math.Ceil(math.Ceil(float64(5*N-Y))/4)/4) + 1)
		}
	} else {
		if int(math.Ceil(math.Trunc(float64(10*N-Y))/9)/4)%5 == Y%5 {
			z = int(math.Ceil(math.Trunc(float64(10*N-Y)) / 9))
		} else {
			z = Y % 5 * (int(math.Ceil(math.Trunc(float64(10*N-Y))/9)/4) - 1)
		}
	}
	if z < 0 || z > MinOf(N, Y) {
		fmt.Println(-1, -1, -1)
		return
	}
	a := N - z
	b := Y - z
	x = (b - 5*a) / 5
	y = N - x - z
	if 10*x+5*y+z == Y {
		fmt.Println(x, y, z)
	} else {
		fmt.Println(-1, -1, -1)
	}
	// fmt.Println(x, y, z)
	// fmt.Println(x+y+z, 10000*x+5000*y+1000*z)
}

func MinOf(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}
