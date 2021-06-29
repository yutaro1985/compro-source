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

type Pair struct {
	v, w int
}

func main() {
	N, W := nextInt(), nextInt()
	Buggage := make([]Pair, N)
	var maxv, maxw, sumv int
	for i := 0; i < N; i++ {
		Buggage[i] = Pair{nextInt(), nextInt()}
		sumv += Buggage[i].v
		ChmaxInt(&maxv, Buggage[i].v)
		ChmaxInt(&maxw, Buggage[i].w)
	}
	var ans int
	if N <= 30 {
		// 半分全列挙
		A, B := make([]Pair, N/2), make([]Pair, (N+1)/2)
		copy(A, Buggage[:N/2])
		copy(B, Buggage[N/2:])
		Alist := make([]Pair, 0)
		for bit := 0; bit < 1<<len(A); bit++ {
			var w, v int
			for i := 0; i < len(A); i++ {
				if bit>>i&1 == 1 {
					w += A[i].w
					v += A[i].v
				}
			}
			Alist = append(Alist, Pair{v, w})
		}
		sort.Slice(Alist, func(i, j int) bool { return Alist[i].w <= Alist[j].w })
		AlistO := make([]Pair, 0)
		AlistO = append(AlistO, Alist[0])
		tmp := Alist[0].v
		for i := 1; i < len(Alist); i++ {
			if Alist[i].v > tmp {
				AlistO = append(AlistO, Alist[i])
				tmp = Alist[i].v
			}
		}
		for bit := 0; bit < 1<<len(B); bit++ {
			var w, v int
			for i := 0; i < len(B); i++ {
				if bit>>i&1 == 1 {
					w += B[i].w
					v += B[i].v
				}
			}
			if w > W {
				continue
			}
			idx := sort.Search(len(AlistO), func(i int) bool {
				return w+AlistO[i].w > W
			})
			idx = MaxOf(0, idx-1)
			ChmaxInt(&ans, v+AlistO[idx].v)
		}
	} else if maxv <= 1000 {
		DP := initInts(sumv+7, INF)
		DP[0] = 0
		for i := 0; i < N; i++ {
			tmp := make([]int, len(DP))
			copy(tmp, DP)
			for j := 0; j < len(DP); j++ {
				if j >= Buggage[i].v {
					ChminInt(&tmp[j], DP[j-Buggage[i].v]+Buggage[i].w)
				}
			}
			DP, tmp = tmp, DP
		}
		for i := 0; i < len(DP); i++ {
			if DP[i] <= W {
				ans = i
			}
		}
	} else {
		DP := initInts(W+1, 0)
		for i := 0; i < N; i++ {
			tmp := make([]int, len(DP))
			copy(tmp, DP)
			for j := 0; j < len(DP); j++ {
				if j >= Buggage[i].w {
					ChmaxInt(&tmp[j], (DP[j-Buggage[i].w] + Buggage[i].v))
				}
			}
			DP, tmp = tmp, DP
		}
		for _, v := range DP {
			ChmaxInt(&ans, v)
		}
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
