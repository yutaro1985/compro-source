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

// 巡回セールスマン問題
// bitDPを使う
// 解説参照
// https://atcoder.jp/contests/abc180/editorial/154
// https://www.youtube.com/watch?v=r4ujcFBDBw4&t=3210s

func main() {
	N := nextInt()
	Cities := make([]City, N)
	distance := make([][]int, N)
	N2 := 1 << N
	dp := make([][]int, N2+1)
	for i := 0; i <= N2; i++ {
		dp[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = INF
		}
	}
	for i := range distance {
		distance[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		Cities[i] = City{nextInt(), nextInt(), nextInt()}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			distance[i][j] = GetCost(Cities[i], Cities[j])
		}
	}
	for i := 0; i < N; i++ {
		if i == 0 {
			continue
		}
		dp[1<<i][i] = distance[0][i]
	}
	for bit := 0; bit < N2; bit++ {
		for i := 0; i < N; i++ {
			// iは遷移元。訪れてない場合は飛ばす
			if bit>>i&1 == 0 {
				continue
			}
			for j := 0; j < N; j++ {
				//jは遷移先。訪問済みだったら飛ばす
				if bit>>j&1 == 1 {
					continue
				}
				dp[bit|1<<j][j] = MinOf(dp[bit|1<<j][j], dp[bit][i]+distance[i][j])
			}
		}
	}
	fmt.Println(dp[N2-1][0])
}

type Pair struct {
	index int
	Cost  int
}

type City struct {
	X int
	Y int
	Z int
}

func GetCost(A, B City) int {
	return Abs(A.X-B.X) + Abs(A.Y-B.Y) + MaxOf(0, B.Z-A.Z)
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

func ceil(a, b int) int {
	return (a + (b - 1)) / b
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
