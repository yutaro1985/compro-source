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

// 見様見真似
// https://drken1215.hatenablog.com/entry/2020/01/26/225000

func main() {
	H, N := nextInt(), nextInt()
	Majics := make([]Majic, N)
	for i := 0; i < N; i++ {
		Majics[i] = Majic{nextInt(), nextInt()}
	}
	DP := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		DP[i] = make([]int, H+1)
		for j := range DP[i] {
			DP[i][j] = INF
		}
	}
	DP[0][0] = 0
	for i := 0; i < N; i++ {
		for j := 0; j <= H; j++ {
			DP[i+1][j] = MinOf(DP[i+1][j], DP[i][j])
			DP[i+1][MinOf(j+Majics[i].A, H)] = MinOf(DP[i+1][MinOf(j+Majics[i].A, H)], DP[i+1][j]+Majics[i].B)
		}
	}
	fmt.Println(DP[N][H])
}

type Majic struct {
	A int
	B int
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
