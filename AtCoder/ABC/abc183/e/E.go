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

// わからなかったので解説参照
// https://atcoder.jp/contests/abc183/editorial/179
// https://www.youtube.com/watch?v=N_Xtv_mctTw&t=2448s

func main() {
	H, W := nextInt(), nextInt()
	S := makeStrings(H)
	DP := make([][]mint, H)
	DPx := make([][]mint, H)
	DPy := make([][]mint, H)
	DPxy := make([][]mint, H)
	for i := 0; i < H; i++ {
		DP[i] = make([]mint, W)
		DPx[i] = make([]mint, W)
		DPy[i] = make([]mint, W)
		DPxy[i] = make([]mint, W)
	}
	DP[0][0] = 1
	DPx[0][0] = 1
	DPy[0][0] = 1
	DPxy[0][0] = 1
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if S[i][j] == '#' {
				continue
			}
			// 横方向の累積和を一つの前の値から求める
			if j > 0 {
				DPx[i][j].addSelf(DPx[i][j-1])
			}
			// 縦方向の累積和を一つの前の値から求める
			if i > 0 {
				DPy[i][j].addSelf(DPy[i-1][j])
			}
			// 斜めの累積和を一つの前の値から求める
			if i > 0 && j > 0 {
				DPxy[i][j].addSelf(DPxy[i-1][j-1])
			}
			// i,jの各方向の累積和の値を足す
			DP[i][j].addSelf(DPx[i][j])
			DP[i][j].addSelf(DPy[i][j])
			DP[i][j].addSelf(DPxy[i][j])
			// i,jの和を求めたら、それを各方向の累積和に反映する
			DPx[i][j].addSelf(DP[i][j])
			DPy[i][j].addSelf(DP[i][j])
			DPxy[i][j].addSelf(DP[i][j])
		}
	}
	fmt.Println(DP[H-1][W-1])
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
	res := mint(1)
	for i := 0; i < K; i++ {
		res.mulSelf(mint(N - i))
		res.divSelf(mint(i + 1))
	}
	return res
}
