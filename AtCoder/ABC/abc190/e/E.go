package main

import (
	"bufio"
	"container/list"
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
	INF     = 1 << 60
)

// 巡回セールスマン問題
// TODO あとでもう一度やってみる

func main() {
	N, M := nextInt(), nextInt()
	G := make([][]int, N)
	for i := 0; i < M; i++ {
		A, B := nextInt()-1, nextInt()-1
		G[A] = append(G[A], B)
		G[B] = append(G[B], A)
	}
	K := nextInt()
	C := make([]int, K)
	for i := 0; i < K; i++ {
		C[i] = nextInt() - 1
	}
	// 見るのは必要なK個の点
	// K個の点だけ見た時に、それぞれの点同士の最短経路を求めてグラフを作る
	dist := init2DInts(K, K, 0)
	bfs := func(p int) []int {
		dist := initInts(N, INF)
		q := list.New()
		q.PushBack(p)
		dist[p] = 0
		for q.Len() > 0 {
			cur := q.Front().Value.(int)
			q.Remove(q.Front())
			for _, next := range G[cur] {
				if dist[next] != INF {
					continue
				}
				dist[next] = dist[cur] + 1
				q.PushBack(next)
			}
		}
		return dist
	}
	for i := 0; i < K; i++ {
		d := bfs(C[i])
		for j := 0; j < K; j++ {
			dist[i][j] = d[C[j]]
		}
	}
	K2 := 1 << K
	DP := init2DInts(K2, K, INF)
	for i := 0; i < K; i++ {
		DP[1<<i][i] = 1
	}
	for bit := 0; bit < K2; bit++ {
		for i := 0; i < K; i++ {
			if bit>>i&1 == 0 {
				continue
			}
			for j := 0; j < K; j++ {
				if bit>>j&1 == 1 {
					continue
				}
				ChminInt(&DP[bit|1<<j][j], DP[bit][i]+dist[i][j])
			}
		}
	}
	ans := INF
	for i := 0; i < K; i++ {
		ChminInt(&ans, DP[K2-1][i])
	}
	if ans == INF {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
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
