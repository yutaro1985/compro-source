package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
// 全体からあるマスを照らせない置き方を引く
// 解説のやり方とは別に、UnionFindで照らせるかどうかをチェックする
// 参考 https://atcoder.jp/contests/hhkb2020/submissions/17334962

// それはそれとして、典型手法らしいのであとで思い出せるように
// https://drken1215.hatenablog.com/entry/2020/10/11/155900

// 類似問題あとでやる
// https://atcoder.jp/contests/abc129/tasks/abc129_d

func main() {
	H, W := nextInt(), nextInt()
	S := make([]string, H)
	var K int
	for i := 0; i < H; i++ {
		S[i] = nextLine()
		K += strings.Count(S[i], ".")
	}
	Pow2 := make([]int, K+1)
	Pow2[0] = 1
	for i := 0; i < K; i++ {
		Pow2[i+1] = (Pow2[i] * 2) % mod
	}
	var ans int
	ufH, ufW := newUnionFind(H*W), newUnionFind(H*W)
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if S[i][j] == '#' {
				continue
			}
			// 各マスにuniqueな値が振られるように計算する
			// 参考にした解法ではiにWをかけた値とjを足した値を使っている。
			if i > 0 && S[i-1][j] == '.' {
				ufH.unite((i-1)*W+j, i*W+j)
			}
			if j > 0 && S[i][j-1] == '.' {
				ufW.unite(i*W+j-1, i*W+j)
			}
		}
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if S[i][j] == '#' {
				continue
			}
			index := i*W + j
			// 重複する1マスを引く
			size := ufH.size(index) + ufW.size(index) - 1
			ans += (Pow2[K] - Pow2[K-size] + mod) % mod
			ans %= mod
		}
	}
	// ans = ((Pow2[K]*K)%mod - unlightenable) % mod
	fmt.Println(ans)
}

// UnionFind の定義
// 下記を参考
// https://youtu.be/TdR816rqc3s?t=6822
// https://github.com/atcoder/live_library/blob/master/uf.cpp
// https://qiita.com/haru1843/items/2295d0ec1f5002bd5c33#%E5%AE%9F%E8%A3%85
type UnionFind struct {
	parent []int
	maxlen int
}

// UnionFind のスライス初期化
func newUnionFind(N int) *UnionFind {
	u := new(UnionFind)
	u.parent = make([]int, N)
	u.maxlen = 1
	for i := range u.parent {
		u.parent[i] = -1
	}
	return u
}

// xの根を見つける
func (u *UnionFind) find(x int) int {
	if u.parent[x] < 0 {
		return x
	}
	u.parent[x] = u.find(u.parent[x])
	return u.parent[x]
}

// xとyのグループを結合する
func (u *UnionFind) unite(x, y int) {
	xf := u.find(x)
	yf := u.find(y)
	if xf == yf {
		return
	}
	// 常に大きい方に小さい木をくっつける
	if xf > yf {
		xf, yf = yf, xf
	}
	u.parent[xf] += u.parent[yf]
	u.parent[yf] = xf
	if u.parent[xf] < 0 {
		u.maxlen = MaxOf(u.maxlen, Abs(u.parent[xf]))
	}
}

// xとyが同じグループに所属するかどうかを返す
func (u *UnionFind) same(x, y int) bool {
	if u.find(x) == u.find(y) {
		return true
	}
	return false
}

// xの所属するグループの木の大きさを返す
func (u *UnionFind) size(x int) int {
	return -u.parent[u.find(x)]
}

// 根の数を返す
func (u UnionFind) rootcnt() int {
	var cnt int
	for _, v := range u.parent {
		if v < 0 {
			cnt++
		}
	}
	return cnt
}

var d = []Position{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

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
