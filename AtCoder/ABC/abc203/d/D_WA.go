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
	N, K := nextInt(), nextInt()
	A := make([][]int, N)
	for i := 0; i < N; i++ {
		A[i] = makeInts(N)
	}
	ans := INF
	Partition := func(lst []int, pivot int) int {
		pivotidx := -1
		for k, v := range lst {
			if v == pivot {
				pivotidx = k
			}
		}
		// if pivotidx == -1 {
		// 	panic(pivotidx)
		// }
		lst[pivotidx], lst[len(lst)-1] = lst[len(lst)-1], lst[pivotidx]
		pivot = lst[len(lst)-1]
		i := -1
		for j, v := range lst[:len(lst)-1] {
			if v <= pivot {
				i++
				lst[i], lst[j] = lst[j], lst[i]
			}
		}
		lst[i+1], lst[len(lst)-1] = lst[len(lst)-1], lst[i+1]
		return i + 1
	}
	var Select func([]int, int) int
	Select = func(lst []int, i int) int {
		if len(lst) == 1 {
			return lst[0]
		}
		splitlst := make([][]int, 1)
		for i := 0; i < len(lst); i++ {
			splitlst[len(splitlst)-1] = append(splitlst[len(splitlst)-1], lst[i])
			if len(splitlst[len(splitlst)-1]) == 5 {
				sort.Ints((splitlst[len(splitlst)-1]))
				splitlst = append(splitlst, []int{})
			}
		}
		meds := make([]int, 0)
		for i := 0; i < len(splitlst); i++ {
			size := len(splitlst[i]) - 1
			meds = append(meds, splitlst[i][size/2])
		}
		x := Select(meds, (len(meds)-1)/2)
		k := Partition(lst, x)
		if i == k {
			return x
		} else if i < k {
			return Select(lst[:k], i)
		} else {
			return Select(lst[k+1:], i-(k+1))
		}
	}
	for i := 0; i+K-1 < N; i++ {
		for j := 0; j+K-1 < N; j++ {
			ponds := make([]int, 0)
			// if i+K >= N {
			// 	continue
			// }
			// if j+K >= N {
			// 	continue
			// }
			for k := i; k < i+K; k++ {
				for l := j; l < j+K; l++ {
					ponds = append(ponds, A[k][l])
				}
			}
			meds := Select(ponds, (len(ponds)-1)/2)
			ChminInt(&ans, meds)
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
