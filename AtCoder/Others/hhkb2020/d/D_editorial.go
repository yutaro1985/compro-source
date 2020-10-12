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
)

// 解説を参照
// youtube解説
// 重ならないような分割線を動かしていって、それぞれについて組み合わせ数を求めて足す
// 重複分は分割線を縦横に引いて、縦でも横でも分割可能なパターンを計算する
// ※こちらはmodintがないと計算がきつい…
// https://youtu.be/eus_giFYAIs?t=5820

func f1(N int) int {
	if N < 0 {
		return 0
	}
	res := (N * (N + 1) / 2) % mod
	return res
}

func f2(N int) int {
	if N < 0 {
		return 0
	}
	res := f1(N)
	res = (res * res) % mod
	return res
}

func solve() {
	N, A, B := nextInt(), nextInt(), nextInt()
	x := (((f1(N-A-B+1) * (N - B + 1)) % mod) * (N - A + 1)) % mod
	y := f2(N - A - B + 1)
	ans := (x*4 - y*4 + mod) % mod
	fmt.Println(ans)
}

func main() {
	T := nextInt()
	for i := 0; i < T; i++ {
		solve()
	}
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
