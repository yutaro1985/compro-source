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
	INF     = 1 << 60
)

type House struct {
	i, x, y int
}

type Pair struct {
	A, B int
}

func main() {
	N := nextInt()
	H := make([]House, N)
	for i := 0; i < N; i++ {
		H[i] = House{i, nextInt(), nextInt()}
	}
	H2 := make([]House, N)
	copy(H2, H)
	sort.Slice(H, func(i, j int) bool {
		return H[i].x < H[j].x
	})
	sort.Slice(H2, func(i, j int) bool {
		return H2[i].y < H2[j].y
	})
	ans := make([]int, 0)
	memo := make(map[Pair]int)
	// N-1 - 0
	memo[Pair{MinOf(H[0].i, H[N-1].i), MaxOf(H[0].i, H[N-1].i)}] = H[N-1].x - H[0].x
	// N-1 -1
	memo[Pair{MinOf(H[1].i, H[N-1].i), MaxOf(H[1].i, H[N-1].i)}] = H[N-1].x - H[1].x
	// N-2 -0
	memo[Pair{MinOf(H[0].i, H[N-2].i), MaxOf(H[0].i, H[N-2].i)}] = H[N-2].x - H[0].x
	// N-1 - 0
	if _, e := memo[Pair{MinOf(H[0].i, H[N-1].i), MaxOf(H[0].i, H[N-1].i)}]; !e {
		memo[Pair{MinOf(H2[0].i, H2[N-1].i), MaxOf(H2[0].i, H2[N-1].i)}] = H2[N-1].y - H2[0].y
	} else if memo[Pair{MinOf(H2[0].i, H2[N-1].i), MaxOf(H2[0].i, H2[N-1].i)}] < H2[N-1].y-H2[0].y {
		memo[Pair{MinOf(H2[0].i, H2[N-1].i), MaxOf(H2[0].i, H2[N-1].i)}] = H2[N-1].y - H2[0].y
	}
	if _, e := memo[Pair{MinOf(H2[1].i, H2[N-1].i), MaxOf(H2[1].i, H2[N-1].i)}]; !e {
		memo[Pair{MinOf(H2[1].i, H2[N-1].i), MaxOf(H2[1].i, H2[N-1].i)}] = H2[N-1].y - H2[1].y
	} else if memo[Pair{MinOf(H2[1].i, H2[N-1].i), MaxOf(H2[1].i, H2[N-1].i)}] < H2[N-1].y-H2[1].y {
		memo[Pair{MinOf(H2[1].i, H2[N-1].i), MaxOf(H2[1].i, H2[N-1].i)}] = H2[N-1].y - H2[1].y
	}
	if _, e := memo[Pair{MinOf(H2[0].i, H2[N-2].i), MaxOf(H2[0].i, H2[N-2].i)}]; !e {
		memo[Pair{MinOf(H2[0].i, H2[N-2].i), MaxOf(H2[0].i, H2[N-2].i)}] = H2[N-2].y - H2[0].y
	} else if memo[Pair{MinOf(H2[0].i, H2[N-2].i), MaxOf(H2[0].i, H2[N-2].i)}] < H2[N-2].y-H2[0].y {
		memo[Pair{MinOf(H2[0].i, H2[N-2].i), MaxOf(H2[0].i, H2[N-2].i)}] = H2[N-2].y - H2[0].y
	}
	for _, v := range memo {
		ans = append(ans, v)
	}
	sort.Ints(ans)
	fmt.Println(ans[len(ans)-2])
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
