package main

import (
	"bufio"
	"fmt"
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

type Bulb struct {
	A, B int
}
type Block struct {
	C, D int
}

type BChecked struct {
	I    int
	H, W bool
}

func index(i, j, w int) int {
	return i*w + j
}

// UnionFindを使ったら実装が重くなりすぎたので解説を参考にやり直す

func main() {
	H, W, N, M := nextInt(), nextInt(), nextInt(), nextInt()
	var ans int
	Bulbs := make([]Bulb, N)
	for i := 0; i < N; i++ {
		Bulbs[i] = Bulb{nextInt() - 1, nextInt() - 1}
	}
	Blocks := make([]Block, M)
	for i := 0; i < M; i++ {
		Blocks[i] = Block{nextInt() - 1, nextInt() - 1}
	}
	S := make([][]byte, H)
	for i := 0; i < H; i++ {
		S[i] = make([]byte, W)
		for j := 0; j < W; j++ {
			S[i][j] = '.'
		}
	}
	Bmemo := make(map[int]bool)
	BChecked := make([]BChecked, H*W)
	for _, v := range Bulbs {
		S[v.A][v.B] = 'B'
		Bmemo[index(v.A, v.B, W)] = true
	}
	for _, v := range Blocks {
		S[v.C][v.D] = '#'
	}

	seen := make([][]bool, H)
	for i := 0; i < H; i++ {
		seen[i] = make([]bool, W)
	}
	sort.Slice(Bulbs, func(i, j int) bool {
		if Bulbs[i].A == Bulbs[j].A {
			return Bulbs[i].B < Bulbs[j].B
		} else {
			return Bulbs[i].A < Bulbs[j].A
		}
	})
	ufH, ufW := newUnionFind(H*W), newUnionFind(H*W)
	for _, b := range Bulbs {
		for i := b.A; i < H-1; i++ {
			if S[i+1][b.B] == '#' {
				break
			}
			if BChecked[index(i+1, b.B, W)].H {
				break
			}
			if ufH.same(index(i, b.B, W), index(i+1, b.B, W)) {
				continue
			}
			ufH.unite(index(i, b.B, W), index(i+1, b.B, W))
			if S[i+1][b.B] == 'B' {
				BChecked[index(i+1, b.B, W)].H = true
			}
		}
		for i := b.A; i > 0; i-- {
			if S[i-1][b.B] == '#' {
				break
			}
			if BChecked[index(i-1, b.B, W)].H {
				break
			}
			if ufH.same(index(i, b.B, W), index(i-1, b.B, W)) {
				continue
			}
			ufH.unite(index(i, b.B, W), index(i-1, b.B, W))
			if S[i-1][b.B] == 'B' {
				BChecked[index(i-1, b.B, W)].H = true
			}
		}
		for i := b.B; i < W-1; i++ {
			if S[b.A][i+1] == '#' {
				break
			}
			if BChecked[index(b.A, i+1, W)].W {
				break
			}
			if ufW.same(index(b.A, i, W), index(b.A, i+1, W)) {
				continue
			}
			ufW.unite(index(b.A, i, W), index(b.A, i+1, W))
			if S[b.A][i+1] == 'B' {
				BChecked[index(b.A, i+1, W)].W = true
			}
		}
		for i := b.B; i > 0; i-- {
			if S[b.A][i-1] == '#' {
				break
			}
			if BChecked[index(b.A, i-1, W)].W {
				break
			}
			if ufW.same(index(b.A, i, W), index(b.A, i-1, W)) {
				continue
			}
			ufW.unite(index(b.A, i, W), index(b.A, i-1, W))
			if S[b.A][i-1] == 'B' {
				BChecked[index(b.A, i-1, W)].W = true
			}
		}
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			Hsize := ufH.size(index(i, j, W))
			if Hsize > 1 {
				seen[i][j] = true
			} else if Hsize == 1 {
				if _, e := Bmemo[index(i, j, W)]; e {
					seen[i][j] = true
				}
			}
			Wsize := ufW.size(index(i, j, W))
			if Wsize > 1 {
				seen[i][j] = true
			} else if Wsize == -1 {
				if _, e := Bmemo[index(i, j, W)]; e {
					seen[i][j] = true
				}
			}
		}
	}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if seen[i][j] {
				ans++
			}
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

// UnionFind の定義
// 下記を参考
// https://youtu.be/TdR816rqc3s?t=6822
// https://github.com/atcoder/live_library/blob/master/uf.cpp
// https://qiita.com/haru1843/items/2295d0ec1f5002bd5c33#%E5%AE%9F%E8%A3%85
type UnionFind struct {
	parent []int
	maxlen int
	N      int
}

// UnionFind のスライス初期化
func newUnionFind(N int) *UnionFind {
	u := new(UnionFind)
	u.N = N
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

func (uf UnionFind) groups() [][]int {
	rootBuf, groupSize := make([]int, uf.N), make([]int, uf.N)
	for i := 0; i < uf.N; i++ {
		rootBuf[i] = uf.find(i)
		groupSize[rootBuf[i]]++
	}
	res := make([][]int, uf.N)
	for i := 0; i < uf.N; i++ {
		res[i] = make([]int, 0, groupSize[i])
	}
	for i := 0; i < uf.N; i++ {
		res[rootBuf[i]] = append(res[rootBuf[i]], i)
	}
	result := make([][]int, 0, uf.N)
	for i := 0; i < uf.N; i++ {
		if len(res[i]) != 0 {
			r := make([]int, len(res[i]))
			copy(r, res[i])
			result = append(result, r)
		}
	}
	return result
}
