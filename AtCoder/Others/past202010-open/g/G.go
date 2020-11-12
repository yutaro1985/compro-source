package main

import (
	"bufio"
	"container/list"
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
	N, M := nextInt(), nextInt()
	S := makeStrings(N)
	wall := make([]Position, 0)

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if S[i][j] == '#' {
				wall = append(wall, Position{i, j})
			}
		}
	}
	var ans int
	for _, w := range wall {
		start := w
		seen := init2DBools(N, M, false)
		Sm := make([][]byte, N)
		for i := 0; i < N; i++ {
			Sm[i] = make([]byte, M)
			for j := 0; j < M; j++ {
				if w.H == i && w.W == j {
					Sm[i][j] = '.'
				} else {
					Sm[i][j] = S[i][j]
				}
			}
		}
		q := list.New()
		q.PushBack(start)
		for q.Len() > 0 {
			p := q.Front().Value.(Position)
			q.Remove(q.Front())
			if Sm[p.H][p.W] == '#' {
				continue
			}
			for _, v := range d {
				nextP := Position{p.H + v.H, p.W + v.W}
				if nextP.H < 0 || nextP.H >= N || nextP.W < 0 || nextP.W >= M {
					continue
				}
				if Sm[nextP.H][nextP.W] == '#' {
					continue
				}
				if !seen[nextP.H][nextP.W] {
					seen[nextP.H][nextP.W] = true
					q.PushBack(nextP)
				}
			}
		}
		// fmt.Println(seen, w)
		isOk := true
		for i := 0; i < N; i++ {
			for j := 0; j < M; j++ {
				if !seen[i][j] && Sm[i][j] == '.' {
					isOk = false
					break
				}
			}
		}
		if isOk {
			ans++
		}
	}
	fmt.Println(ans)
}

var d = []Position{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var d8 = []Position{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

// Position として迷路問題での現在地を表す構造体を定義
type Position struct {
	H int
	W int
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

func makeInts(N int) []int {
	res := make([]int, N)
	for i := range res {
		res[i] = nextInt()
	}
	return res
}

func makeStrings(N int) []string {
	res := make([]string, N)
	for i := range res {
		res[i] = nextLine()
	}
	return res
}

func init2DInts(H, W, N int) [][]int {
	res := make([][]int, H)
	for i := 0; i < H; i++ {
		res[i] = make([]int, W)
		if N != 0 {
			for j := 0; j < W; j++ {
				res[i][j] = N
			}
		}
	}
	return res
}

func init2DBools(H, W int, b bool) [][]bool {
	res := make([][]bool, H)
	for i := 0; i < H; i++ {
		res[i] = make([]bool, W)
		if b {
			for j := 0; j < W; j++ {
				res[i][j] = true
			}
		}
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
