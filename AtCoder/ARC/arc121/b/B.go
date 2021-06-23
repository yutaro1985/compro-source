package main

import (
	"bufio"
	"fmt"
	"math"
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

type Dog struct {
	a int
	c string
}

type Pair struct {
	A, B int
}

func main() {
	N := nextInt()
	Dogs := make([]Dog, N*2)
	for i := 0; i < N*2; i++ {
		Dogs[i] = Dog{nextInt(), nextLine()}
	}
	R := make([]int, 0)
	G := make([]int, 0)
	B := make([]int, 0)
	for i := 0; i < len(Dogs); i++ {
		switch Dogs[i].c {
		case "R":
			R = append(R, Dogs[i].a)
		case "G":
			G = append(G, Dogs[i].a)
		case "B":
			B = append(B, Dogs[i].a)
		}
	}
	sort.Ints(R)
	sort.Ints(G)
	sort.Ints(B)
	// fmt.Println(len(R), len(G), len(B))
	ans := math.MaxInt64
	if len(R)%2 == 0 && len(G)%2 == 0 && len(B)%2 == 0 {
		fmt.Println(0)
		return
	}

	dist2 := func(M, N []int) int {
		res := math.MaxInt64
		if len(M) < len(N) {
			M, N = N, M
		}
		if len(N) == 0 {
			if len(M) == 0 {
				return 0
			} else {
				return M[0]
			}
		}

		for i := 0; i < len(M); i++ {
			idx := sort.SearchInts(N, M[i])
			for j := idx - 10; j <= idx+10; j++ {
				if j < 0 || j >= len(N) {
					continue
				}
				ChminInt(&res, Abs(M[i]-N[j]))
			}
		}
		return res
	}
	// dist := func(M, N, L []int) int {
	// 	res := math.MaxInt64
	// 	ChminInt(&res, dist2(M, N))
	// 	NL := dist2(N, L)
	// 	LM := dist2(L, M)
	// 	ChminInt(&res, NL+LM)
	// 	return res
	// }
	if len(R)%2 == 0 {
		// ChminInt(&ans, dist(G, B, R))
		ChminInt(&ans, MinOf(dist2(G, B), dist2(R, G)+dist2(B, R)))
	} else if len(G)%2 == 0 {
		// ChminInt(&ans, dist(B, R, G))
		ChminInt(&ans, MinOf(dist2(B, R), dist2(G, B)+dist2(R, G)))
	} else if len(B)%2 == 0 {
		// ChminInt(&ans, dist(R, G, B))
		ChminInt(&ans, MinOf(dist2(R, G), dist2(G, B)+dist2(B, R)))
	}
	fmt.Println(ans)
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
