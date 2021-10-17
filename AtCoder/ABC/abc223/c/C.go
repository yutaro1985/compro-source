package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
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

type Point struct {
	A, B float64
}

func main() {
	N := nextInt()
	P := make([]Point, N)
	Csum := make([]float64, N+1)
	CsumR := make([]float64, N+1)
	for i := 0; i < N; i++ {
		P[i] = Point{nextFloat64(), nextFloat64()}
		Csum[i+1] = Csum[i] + P[i].A/P[i].B
	}
	for i := 0; i < N; i++ {
		CsumR[N-1-i] = CsumR[N-i] + P[N-1-i].A/P[N-1-i].B
	}
	var idx int
	var ans float64
	for i := 0; i <= N; i++ {
		A := big.NewFloat(Csum[i])
		B := big.NewFloat(CsumR[i])
		// LocalPrint(A, B, ans)
		if A.Cmp(B) >= 0 {
			idx = i
			break
		}
	}
	for i := 0; i < idx-1; i++ {
		ans += P[i].A
	}
	if Csum[idx] == CsumR[idx] {
		fmt.Println(ans + P[idx-1].A)
		return
	}
	// LocalPrint(Csum)
	// LocalPrint(CsumR)
	// LocalPrint(idx)
	// LocalPrint(P[idx-1].A)
	// LocalPrint("ans", ans)
	r := P[idx-1].A
	v := P[idx-1].B
	// LocalPrint(r, v)
	// LocalPrint(Csum[idx-1], CsumR[idx], Csum[idx-1]-CsumR[idx], r-math.Abs((Csum[idx-1]-CsumR[idx])*v))
	r -= v * math.Abs(Csum[idx-1]-CsumR[idx])
	// LocalPrint(r)
	A := big.NewFloat(Csum[idx-1])
	B := big.NewFloat(CsumR[idx])
	if A.Cmp(B) >= 0 {
		ans += r / (v * 2) * v
	} else {
		ans += P[idx-1].A - r/(v*2)*v
	}
	// ans += ((r - v*math.Abs(Csum[idx]-CsumR[idx])) / (v * 2)) * v
	// ans += (P[idx-1].A - P[idx-1].B*math.Abs(Csum[idx]-CsumR[idx])) / (P[idx-1].B * 2) * P[idx-1].B
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
