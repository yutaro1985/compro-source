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

func main() {
	N, Q := nextInt(), nextInt()
	follows := make([][]bool, N)
	// followers := make([][]bool, N)
	ans := make([][]byte, N)
	for i := 0; i < N; i++ {
		follows[i] = make([]bool, N)
		// followers[i] = make([]bool, N)
		ans[i] = make([]byte, N)
	}
	for i := 0; i < Q; i++ {
		ope := nextInt()
		switch ope {
		case 1:
			a, b := nextInt()-1, nextInt()-1
			follows[a][b] = true
			// followers[b][a] = true
		case 2:
			a := nextInt() - 1
			for j := 0; j < N; j++ {
				if follows[j][a] {
					follows[a][j] = true
				}
			}
		case 3:
			a := nextInt() - 1
			addfollow := make([]int, 0)
			for j, v := range follows[a] {
				if v {
					addfollow = append(addfollow, j)
				}
			}
			for _, v := range addfollow {
				for j, vf := range follows[v] {
					if vf {
						follows[a][j] = true
					}
				}
			}
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			switch follows[i][j] && i != j {
			case true:
				ans[i][j] = 'Y'
			case false:
				ans[i][j] = 'N'
			}
		}
	}
	for i := 0; i < N; i++ {
		fmt.Println(string(ans[i]))
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

// ChminInt 第一引数のほうが大きかった場合第二引数の値を代入する。
// 1つ目の値は参照渡しする
func ChminInt(a *int, b int) {
	if *a > b {
		*a = b
	}
}

// ChmanInt 第一引数のほうが小さかった場合第二引数の値を代入する。
// 1つ目の値は参照渡しする
func ChmanInt(a *int, b int) {
	if *a < b {
		*a = b
	}
}
