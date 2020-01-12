package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		n, _ := strconv.Atoi(sc.Text())
		if n%2 == 0 {
			fmt.Println("0")
		} else {
			fmt.Println(math.Ceil(float64((1999999 / n)) / 2.0))
		}
		// 隣り合う二値を合計すると3,5,7,9,11…1999999と、奇数が並ぶ
		// その為偶数の倍数は存在しない
		// また、1999999をNで割ると、1~1999999の中で割り切れる数の個数が
		// 偶数奇数含めて求まる。
		// 奇数を2倍していくと奇数・偶数・奇数・偶数…と続くため
		// 割り切れる数の個数のうち、半分、2で割り切れない場合は小数点を切り上げた値が
		// 3,5,～1999999に含まれるNの倍数の数となる。
	}
}
