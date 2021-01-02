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
	INF     = 1 << 60
)

var N, Q int
var G [][]int
var A, B []int
var depth []int
var dp []int

// 解説放送での解法
// こちらの問題が類題
// https://drken1215.hatenablog.com/entry/2020/05/18/174000

func main() {
	N = nextInt()
	G = make([][]int, N)
	A = make([]int, N-1)
	B = make([]int, N-1)
	depth = make([]int, N)
	dp = make([]int, N)
	for i := 0; i < N; i++ {
		depth[i] = -1
	}
	for i := 0; i < N-1; i++ {
		a, b := nextInt()-1, nextInt()-1
		A[i] = a
		B[i] = b
		G[a] = append(G[a], b)
		G[b] = append(G[b], a)
	}
	getDepth(0, 0)
	Q = nextInt()
	for i := 0; i < Q; i++ {
		t, e, x := nextInt(), nextInt()-1, nextInt()
		var va, vb int
		// vaは通る点、vbは通らない点
		switch t {
		case 1:
			va = A[e]
			vb = B[e]
		case 2:
			va = B[e]
			vb = A[e]
		}
		// 通る点が上の場合、全体に足して部分木以下は引く
		// depthの値が小さい方が上
		if depth[va] < depth[vb] {
			dp[0] += x
			dp[vb] -= x
		} else {
			// そうでない場合、部分木だけに足す
			dp[va] += x
		}
	}
	dfs(0, 0)
	for _, v := range dp {
		fmt.Println(v)
	}
}

func getDepth(v, p int) {
	depth[v] = p
	for _, next := range G[v] {
		if depth[next] == -1 {
			getDepth(next, p+1)
		}
	}
}

func dfs(a, now int) {
	now += dp[a]
	dp[a] = now
	for _, next := range G[a] {
		if depth[next] > depth[a] {
			dfs(next, now)
		}
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
