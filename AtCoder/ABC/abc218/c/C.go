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
	mod     = int(1e9) + 7
	maxsize = 510000
	INF     = 1 << 60
)

func main() {
	N := nextInt()
	var cnt1, cnt2 int
	S := makeStrings(N)
	T := makeStrings(N)
	for i := 0; i < N; i++ {
		cnt1 += strings.Count(S[i], "#")
		cnt2 += strings.Count(T[i], "#")
	}
	if cnt1 != cnt2 {
		fmt.Println("No")
		return
	}
	Solve := func(S, T []string) bool {
		var l, r, t, b int
		for i := 0; i < N; i++ {
			isOK := false
			for j := 0; j < N; j++ {
				if T[j][i] == '#' {
					l = i
					isOK = true
					break
				}
			}
			if isOK {
				break
			}
		}
		for i := 0; i < N; i++ {
			isOK := false
			for j := 0; j < N; j++ {
				if T[i][j] == '#' {
					t = i
					isOK = true
					break
				}
			}
			if isOK {
				break
			}
		}
		for i := N - 1; i >= 0; i-- {
			isOK := false
			for j := 0; j < N; j++ {
				if T[i][j] == '#' {
					b = i
					isOK = true
					break
				}
			}
			if isOK {
				break
			}
		}
		for i := N - 1; i >= 0; i-- {
			isOK := false
			for j := 0; j < N; j++ {
				if T[j][i] == '#' {
					r = i
					isOK = true
					break
				}
			}
			if isOK {
				break
			}
		}
		// LocalPrint(l, r, t, b)
		check := make([][]byte, b-t+1)
		for i := 0; i <= b-t; i++ {
			check[i] = make([]byte, r-l+1)
			for j := 0; j <= r-l; j++ {
				check[i][j] = T[t+i][l+j]
			}
		}
		// LocalPrint(len(check))
		// for _, v := range check {
		// 	LocalPrint(string(v), len(v))
		// }
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				isOK := true
				for k := 0; k < len(check); k++ {
					if i+k >= N {
						isOK = false
						break
					}
					for l := 0; l < len(check[k]); l++ {
						if j+l >= N {
							isOK = false
							break
						}
						if S[i+k][j+l] != check[k][l] {
							isOK = false
							break
						}
					}
					if !isOK {
						break
					}
				}
				if isOK {
					return true
				}
			}
		}
		return false
	}
	Rotate := func(S []string) []string {
		T := make([][]byte, N)
		for i := 0; i < N; i++ {
			T[i] = make([]byte, N)
		}
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				T[j][N-1-i] = S[i][j]
			}
		}
		res := make([]string, N)
		for i := 0; i < N; i++ {
			res[i] = string(T[i])
		}
		return res
	}
	// 回転なし
	if Solve(S, T) {
		fmt.Println("Yes")
		return
	}
	// 90度回転
	T = Rotate(T)
	if Solve(S, T) {
		fmt.Println("Yes")
		return
	}
	// 180度回転
	T = Rotate(T)
	if Solve(S, T) {
		fmt.Println("Yes")
		return
	}
	// 270度回転
	T = Rotate(T)
	if Solve(S, T) {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
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
