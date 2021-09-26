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
	A := makeInts(N)
	DP := make([][10]mint, N+1)
	DP[0][A[0]] = 1
	for i := 0; i+1 < N; i++ {
		for j := 0; j < 10; j++ {
			if DP[i][j] == 0 {
				continue
			}
			DP[i+1][(j+A[i+1])%10].addSelf(DP[i][j])
			DP[i+1][(j*A[i+1])%10].addSelf(DP[i][j])
		}
	}
	for _, v := range DP[N-1] {
		fmt.Println(v)
	}
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

// SingleCOM はmodintを使ってO(K)で一つの二項係数を求める関数
func SingleCOM(N, K int) mint {
	if N < 0 || K < 0 || K > N {
		return 0
	}
	res := mint(1)
	for i := 0; i < K; i++ {
		res.mulSelf(mint(N - i))
		res.divSelf(mint(i + 1))
	}
	return res
}

// modCOMで使うスライス
var fac = make([]mint, maxsize)
var finv = make([]mint, maxsize)
var inv = make([]mint, maxsize)

// COMinit で COMで使うためのテーブルを作る前処理
// O(N)
// https://qiita.com/drken/items/3b4fdf0a78e7a138cd9a#5-%E4%BA%8C%E9%A0%85%E4%BF%82%E6%95%B0-ncr
func ModCOMinit() {
	fac[0], fac[1] = 1, 1
	finv[0], finv[1] = 1, 1
	inv[1] = 1
	for i := 2; i < maxsize; i++ {
		fac[i] = fac[i-1].mul(mint(i))
		inv[i] = mint(mod) - inv[mod%i].mul(mint(mod/i))
		finv[i] = finv[i-1].mul(inv[i])
	}
}

// FascModCOM nCkを求める。COMinitを先に実行する
// COMinitの結果を使ってO(1)で行える
func FastModCOM(n, k int) mint {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	return fac[n].mul(finv[k].mul(finv[n-k]))
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

func nextFloat64() float64 {
	sc.Scan()
	f, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return f
}

func nextComplex128() complex128 {
	return complex(nextFloat64(), nextFloat64())
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

// LocalPrint はisLocalが1のときだけ標準出力するデバッグ用関数
func LocalPrint(i ...interface{}) {
	if os.Getenv("isLocal") == "1" {
		fmt.Printf("%s", "Local: ")
		fmt.Println(i...)
	}
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
