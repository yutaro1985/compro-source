package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
)

func main() {
	N, M := nextInt(), nextInt()
	H := make([]int, N)
	for i := 0; i < N; i++ {
		H[i] = nextInt()
	}
	sort.Ints(H)
	W := make([]int, M)
	for i := 0; i < M; i++ {
		W[i] = nextInt()
	}
	sort.Ints(W)
	CsumE := make([]int, N+1)
	CsumO := make([]int, N+1)
	for i := 2; i <= N; i++ {
		if i%2 == 0 {
			CsumE[i] = CsumE[i-2] + H[i-1] - H[i-2]
		} else {
			CsumO[i] = CsumO[i-2] + H[i-1] - H[i-2]
		}
	}
	var ans, index int
	minPair := make([]int, N)
	for i := 0; i < N; i++ {
		sub := int(1e9)
		for {
			if index < M-1 {
				if Abs(H[i]-W[index]) < Abs(H[i]-W[index+1]) {
					sub = Abs(H[i] - W[index])
					break
				} else {
					if sub < Abs(H[i]-W[index+1]) {
						break
					} else {
						index++
						sub = Abs(H[i] - W[index])
					}
				}
			} else {
				sub = Abs(H[i] - W[index])
				break
			}
		}
		minPair[i] = sub
	}
	ans = math.MaxInt64 / 100
	for i := 0; i < N; i += 2 {
		var score, front, back int
		if i >= 1 {
			front = CsumE[i] - CsumE[0]
		}
		back = CsumO[N] - CsumO[i+1]
		score = minPair[i] + front + back
		ChminInt(&ans, score)
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

// ChminInt 第一引数のほうが大きかった場合第二引数の値を代入する。
// 1つ目の値は参照渡しする
func ChminInt(a *int, b int) {
	if *a > b {
		*a = b
	}
}

// ChmanInt 第一引数のほうが小さかった場合第二引数の値を代入する。
// 1つ目の値は参照渡しする
func ChmaxInt(a *int, b int) {
	if *a < b {
		*a = b
	}
}
