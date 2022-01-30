package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
	mod = int(1e9) + 7
	// mod     = 998244353
	maxsize = 510000
	INF     = 1 << 60
)

func main() {
	S := nextLine()
	N := len(S)
	var x, y int
	for i := 0; i < N; i++ {
		if S[i] == 'a' {
			x++
		} else {
			break
		}
	}
	for i := N - 1; i >= 0; i-- {
		if S[i] == 'a' {
			y++
		} else {
			break
		}
	}
	// 全てaの場合
	if x == N {
		fmt.Println("Yes")
		return
	}
	// 左のaの数のほうが既に多い場合
	if x > y {
		fmt.Println("No")
		return
	}
	for i := x; i < N-y; i++ {
		if S[i] != S[x-y+N-i-1] {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func preProcess(s string) string {
	var sb strings.Builder
	sb.WriteRune('#')
	for _, c := range s {
		sb.WriteRune(c)
		sb.WriteRune('#')
	}
	return sb.String()
}

func nextProcess(s string) string {
	var sb strings.Builder
	for _, c := range s {
		if c != '#' {
			sb.WriteRune(c)
		}
	}
	return sb.String()
}

func longestPalindrome(s string) string {
	//
	process := preProcess(s)
	//
	p := make([]int, len(process))
	//
	id, mx := 0, 0
	//
	maxLen, maxLenCenter := 0, 0
	//
	for i := range p {
		//
		if i < mx {
			p[i] = min(p[2*id-i], mx-i)
		} else {
			p[i] = 1
		}
		//
		for i-p[i] >= 0 && i+p[i] < len(process) && process[i-p[i]] == process[i+p[i]] {
			p[i]++
		}
		//
		if p[i]+i-1 > mx {
			mx = p[i] + i - 1
			id = i
		}
		//
		if maxLen < p[i]-1 {
			maxLen = p[i] - 1
			maxLenCenter = i
		}
	}
	//
	return nextProcess(process[maxLenCenter-maxLen : maxLenCenter+maxLen])
}

// ReverseString は文字列を反転させたものを返す
func ReverseString(S string) string {
	res := []rune(S)
	for i, j := 0, len(S)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
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
