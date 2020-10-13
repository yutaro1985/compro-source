package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 1e4
	maxBufSize     = 1e6 + 7
)

var buf []byte = make([]byte, maxBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

func main() {
	S := nextLine()
	var ans int
	for bit := 0; bit < 1<<(len(S)-1); bit++ {
		formula := make([]byte, 1)
		formula[0] = S[0]
		for i := 1; i < len(S); i++ {
			// fmt.Println("bit", bit&(1<<(i-1)))
			// ↓この場合計算結果が0,1のみではないので 0か否かを基準にする
			if bit&(1<<(i-1)) != 0 {
				// if bit>>(i-1)&1 == 1 {
				formula = append(formula, '+')
			}
			formula = append(formula, S[i])
		}
		// fmt.Println(string(formula))
		var tmp int
		for i := 0; i < len(formula); i++ {
			if formula[i] == '+' {
				ans += tmp
				tmp = 0
			} else {
				tmp *= 10
				tmp += int(formula[i] - '0')
				if i == len(formula)-1 {
					ans += tmp
				}
			}
		}
	}
	fmt.Println(ans)
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

func pow(p, q int) int {
	res := p
	if q == 0 {
		return 1
	}
	for i := 0; i < q-1; i++ {
		res *= p
	}
	return res
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

func MaxOf(vars ...int) int {
	max := vars[0]
	for _, i := range vars {
		if max < i {
			max = i
		}
	}
	return max
}
