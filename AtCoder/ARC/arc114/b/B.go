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
	mod     = 998244353
	maxsize = 510000
	INF     = 1 << 60
)

func main() {
	N := nextInt()
	f := makeInts(N)
	uf := newUnionFind(N)
	for i := 0; i < N; i++ {
		uf.unite(i, f[i]-1)
	}
	g := uf.groups()
	var ans mint
	ans = mint(PowMod(2, len(g), mod)).sub(1)
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
	if u.parent[xf] >= u.parent[yf] {
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

// UniouFindを各グループに分けたものを2次元スライスで返す
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

type mint int64

func (mi mint) mod() mint {
	mi %= mint(mod)
	if mi < 0 {
		mi += mint(mod)
	}
	return mi
}
func (mi mint) inv() mint {
	// https://qiita.com/drken/items/3b4fdf0a78e7a138cd9a#3-5-%E6%8B%A1%E5%BC%B5-euclid-%E3%81%AE%E4%BA%92%E9%99%A4%E6%B3%95%E3%81%AB%E3%82%88%E3%82%8B%E9%80%86%E5%85%83%E8%A8%88%E7%AE%97
	b, u, v := mint(mod), mint(1), mint(0)
	for b != 0 {
		t := mi / b
		mi -= t * b
		mi, b = b, mi
		u -= t * v
		u, v = v, u
	}
	u %= mint(mod)
	if u < 0 {
		u += mint(mod)
	}
	return u
	// return mi.pow(mint(0).sub(2))
}
func (mi mint) pow(n mint) mint {
	p := mint(1)
	for n > 0 {
		if n&1 == 1 {
			p.mulSelf(mi)
		}
		mi.mulSelf(mi)
		n >>= 1
	}
	return p
}
func (mi mint) factorial() mint {
	// var n mint
	if mi < 0 {
		fmt.Println("[ERROR] Input nums 0 or greater.")
		os.Exit(1)
	}
	if mi == 1 || mi == 0 {
		return 1
	}
	return (mi - 1).factorial().mul(mi)
}
func (mi mint) add(n mint) mint {
	return (mi + n).mod()
}
func (mi mint) sub(n mint) mint {
	return (mi - n).mod()
}
func (mi mint) mul(n mint) mint {
	return (mi * n).mod()
}
func (mi mint) div(n mint) mint {
	return mi.mul(n.inv())
}
func (mi *mint) addSelf(n mint) *mint {
	*mi = mi.add(n)
	return mi
}
func (mi *mint) subSelf(n mint) *mint {
	*mi = mi.sub(n)
	return mi
}
func (mi *mint) mulSelf(n mint) *mint {
	*mi = mi.mul(n)
	return mi
}
func (mi *mint) divSelf(n mint) *mint {
	*mi = mi.div(n)
	return mi
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

func initInts(N, I int) []int {
	res := make([]int, N)
	for i := 0; i < N; i++ {
		res[i] = I
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

// ReverseSort はsort.Interfaceの形を渡すと逆順にソートする
func ReverseSort(data sort.Interface) {
	sort.Sort(sort.Reverse(data))
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
