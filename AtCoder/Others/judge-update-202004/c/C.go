package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
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
	mod = int(1e9) + 7
	// mod     = 998244353
	maxsize = 510000
	INF     = 1 << 60
)

func main() {
	A, B, C := nextInt(), nextInt(), nextInt()
	N := A + B + C
	nums := make([]int, N)
	for i := 0; i < N; i++ {
		nums[i] = i + 1
	}

	var ans int
	for {
		BuildingBlocks := make([][]int, 3)
		for i := 0; i < N; i++ {
			switch {
			case i < A:
				BuildingBlocks[0] = append(BuildingBlocks[0], nums[i])
			case i < A+B:
				BuildingBlocks[1] = append(BuildingBlocks[1], nums[i])
			default:
				BuildingBlocks[2] = append(BuildingBlocks[2], nums[i])
			}
		}
		isOK := true
		for i := 0; i < 3; i++ {
			if i >= 1 {
				for j := 0; j < MinOf(len(BuildingBlocks[i]), len(BuildingBlocks[i-1])); j++ {
					if BuildingBlocks[i][j] < BuildingBlocks[i-1][j] {
						isOK = false
						break
					}
				}
				if !isOK {
					break
				}
			}
			for j := 0; j < len(BuildingBlocks[i])-1; j++ {
				if BuildingBlocks[i][j+1] < BuildingBlocks[i][j] {
					isOK = false
					break
				}
			}
			if !isOK {
				break
			}
		}
		if isOK {
			ans++
		}
		if !NextPermutation(sort.IntSlice(nums)) {
			break
		}
	}
	fmt.Println(ans)
}

// NextPermutation はsortインターフェース（要はソート済みスライス）から次の順列にスライスを並べ替える
// https://play.golang.org/p/ljft9xhOEn
func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
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

// ReverseSlice はSliceを逆順に並べ替える
func ReverseSlice(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
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
