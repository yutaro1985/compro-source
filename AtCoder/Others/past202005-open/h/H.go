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

var buf []byte = make([]byte, initialBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

// 問題によって値は調整する
const (
	mod     = int(1e9) + 7
	maxsize = 510000
	INF     = int(1e9) + 7
)

func main() {
	N, L := nextInt(), nextInt()
	x := make(map[int]bool)
	for i := 0; i < N; i++ {
		x[nextInt()] = true
	}
	T := newInts(3)
	dp := make([]int, L+10)
	for i := 0; i < len(dp); i++ {
		dp[i] = INF
	}
	dp[0] = 0
	for i := 0; i <= L; i++ {
		// 行動1
		if _, e := x[i+1]; e {
			ChminInt(&dp[i+1], dp[i]+T[0]+T[2])
		} else {
			ChminInt(&dp[i+1], dp[i]+T[0])
		}
		// 行動2
		// ChminInt(&dp[i+1], dp[i]+T[0]/2+T[1]/2)
		if _, e := x[i+2]; e {
			ChminInt(&dp[i+2], dp[i]+T[0]+T[1]+T[2])
		} else {
			ChminInt(&dp[i+2], dp[i]+T[0]+T[1])
		}
		// 行動3
		if _, e := x[i+4]; e {
			ChminInt(&dp[i+4], dp[i]+T[0]+T[1]*3+T[2])
		} else {
			ChminInt(&dp[i+4], dp[i]+T[0]+T[1]*3)
		}
	}
	// ジャンプ中にゴールに到達するケース
	ChminInt(&dp[L], dp[L-1]+T[0]/2+T[1]/2)
	ChminInt(&dp[L], dp[L-2]+T[0]/2+T[1]+T[1]/2)
	if L >= 3 {
		ChminInt(&dp[L], dp[L-3]+T[0]/2+T[1]*2+T[1]/2)
	}
	fmt.Println(dp[L])
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

func newInts(N int) []int {
	res := make([]int, N)
	for i := range res {
		res[i] = nextInt()
	}
	return res
}

func newStrings(N int) []string {
	res := make([]string, N)
	for i := range res {
		res[i] = nextLine()
	}
	return res
}

// Math Utilities
// https://play.golang.org/p/bm7uZi0zCN

// Abs Absolute Value
func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

// Pow Integer power: compute a**b using binary powering algorithm
// See Donald Knuth, The Art of Computer Programming, Volume 2, Section 4.6.3
func Pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

// PowMod Modular integer power: compute a**b mod m using binary powering algorithm
func PowMod(a, b, m int) int {
	a = a % m
	p := 1 % m
	for b > 0 {
		if b&1 != 0 {
			p = (p * a) % m
		}
		b >>= 1
		a = (a * a) % m
	}
	return p
}

// Ceil a/bを切り上げたものを返す
func Ceil(a, b int) int {
	return (a + (b - 1)) / b
}

// Mod 負の場合を考慮してあまりを取る
func Mod(a, m int) int {
	res := a % m
	if res < 0 {
		res += m
	}
	return res
}

// MinOf 与えられたintのうち最小のものを返す
func MinOf(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}

// MaxOf 与えられたintのうち最大のものを返す
func MaxOf(vars ...int) int {
	max := vars[0]
	for _, i := range vars {
		if max < i {
			max = i
		}
	}
	return max
}

// ChminInt 第一引数のほうが大きかった場合第二引数の値を代入する。
// 1つ目の値は参照渡しする
func ChminInt(a *int, b int) {
	if *a > b {
		*a = b
	}
}

// ChmaxInt 第一引数のほうが小さかった場合第二引数の値を代入する。
// 1つ目の値は参照渡しする
func ChmaxInt(a *int, b int) {
	if *a < b {
		*a = b
	}
}
