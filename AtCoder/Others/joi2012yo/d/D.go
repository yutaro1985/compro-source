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
	mod     = 10000
	maxsize = 510000
	INF     = 1 << 60
)

// 計算が合わない。また今度

func main() {
	N, K := nextInt(), nextInt()
	P := make(map[int]int)
	for i := 0; i < K; i++ {
		A, B := nextInt(), nextInt()-1
		P[A] = B
	}
	DP := make([][][]mint, N+1)
	for i := 0; i <= N; i++ {
		DP[i] = make([][]mint, 3)
		for j := 0; j < 3; j++ {
			DP[i][j] = make([]mint, 3)
		}
	}
	DP[0][0][0] = 1
	DP[0][0][1] = 0
	DP[0][0][2] = 0
	for i := 0; i < N; i++ {
		p1, e1 := P[i+1]
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if e1 {
					if j == p1 && k == p1 && i >= 2 {
						continue
					}
					DP[i+1][p1][j].addSelf(DP[i][j][k])
				} else {
					for l := 0; l < 3; l++ {
						if l == j && l == k && i >= 2 {
							continue
						}
						DP[i+1][l][j].addSelf(DP[i][j][k])
					}
				}
			}
		}
	}
	var ans mint
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			ans.addSelf(DP[N][i][j])
		}
	}
	fmt.Println(ans)
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
	res := mint(1)
	for i := 0; i < K; i++ {
		res.mulSelf(mint(N - i))
		res.divSelf(mint(i + 1))
	}
	return res
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
