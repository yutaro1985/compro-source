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
	INF     = int(1e9) + 7
)

func main() {
	S := Position{250, 250}
	N, X, Y := nextInt(), nextInt()+S.H, nextInt()+S.W
	H, W := 500, 500
	dist := make([][]int, H)
	for i := 0; i < H; i++ {
		dist[i] = make([]int, W)
		for j := 0; j < W; j++ {
			dist[i][j] = INF
		}
	}
	dist[S.H][S.W] = 0
	for i := 0; i < N; i++ {
		x, y := nextInt()+S.H, nextInt()+S.W
		dist[x][y] = -1
	}
	q := list.New()
	q.PushBack(S)
	for q.Len() > 0 {
		p := q.Front().Value.(Position)
		q.Remove(q.Front())
		for _, v := range d6 {
			nextP := Position{p.H + v.H, p.W + v.W}
			if nextP.H < 0 || nextP.H >= H || nextP.W < 0 || nextP.W >= W {
				continue
			}
			if dist[nextP.H][nextP.W] == -1 {
				continue
			}
			if dist[nextP.H][nextP.W] == INF {
				dist[nextP.H][nextP.W] = dist[p.H][p.W] + 1
				q.PushBack(Position{nextP.H, nextP.W})
			}
		}
	}
	if dist[X][Y] == INF {
		fmt.Println(-1)
	} else {
		fmt.Println(dist[X][Y])
	}
}

var d = []Position{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var d6 = []Position{{1, 0}, {-1, 0}, {0, 1}, {0, -1}, {1, 1}, {-1, 1}}

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
