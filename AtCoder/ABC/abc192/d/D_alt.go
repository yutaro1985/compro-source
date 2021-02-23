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

func main() {
	X := nextLine()
	M := nextInt()
	Xi, _ := strconv.Atoi(X)
	if len(X) == 1 {
		if Xi > M {
			fmt.Println(0)
		} else {
			fmt.Println(1)
		}
		return
	}
	// Ms := strconv.Itoa(M)
	var d int
	for _, v := range X {
		ChmaxInt(&d, int(v-'0'))
	}
	// n := strings.Split(X, "")

	pb := func(a, b, M int) int {
		res := 1
		for ; b > 0; b-- {
			if res*a >= M {
				res = M + 1
				break
			}
			res *= a
		}
		return res
	}
	num := sort.Search(M*2, func(i int) bool {
		if i <= d {
			return false
		}
		var v int
		for j := len(X) - 1; j >= 0; j-- {
			a := pb(i, len(X)-1-j, M)
			x := int(X[j] - '0')
			if v+a*x > M {
				return true
			}
			v += a * x
		}
		// 以下は解説の方法
		// for j := 0; j < len(X); j++ {
		// 	if v > M/i {
		// 		v = M + 1
		// 	} else {
		// 		v = v*i + int(X[j]-'0')
		// 	}
		// 	fmt.Println("v", v)
		// }
		if v > M {
			return true
		} else {
			return false
		}
	})
	num--
	fmt.Println(MaxOf(num-d, 0))
}

func toBase(N string, base uint64) []string {
	// var res uint64
	resS := make([]string, 0)
	Ni, _ := strconv.Atoi(N)
	for Ni > 0 {
		resS = append(resS, strconv.Itoa(Ni%int(base)))
		// resS += strconv.Itoa(Ni % int(base))
		Ni /= int(base)
	}
	rev := func(s []string) {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
	}
	// resS = rev(resS)
	// b := uint64(1)
	// for i := len(N) - 1; i >= 0; i-- {
	// 	res += b * uint64(N[i]-'0')
	// 	b *= base
	// }
	// resS = strconv.FormatUint(res, 10)
	rev(resS)
	return resS
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
