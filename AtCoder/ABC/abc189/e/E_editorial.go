package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
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

type Axis struct {
	A, B int
}

// 行列の積
// https://qiita.com/hamadu/items/fce4ee1e4b5c2c2d24df
// https://gist.github.com/hamadu/193959ef65a660bc1ce3
type Matrix [][]int

// 並列度を指定する
// runtime.GOMAXPROCS(2)

func computePart(i int, a, b, c *Matrix, ch chan int) {
	ac := len((*a)[0])
	bc := len((*b)[0])
	for j := 0; j < bc; j++ {
		part := 0
		for k := 0; k < ac; k++ {
			part += (*a)[i][k] * (*b)[k][j]
		}
		(*c)[i][j] = part
	}
	ch <- 1
}

func mulConcurrent(a, b *Matrix) Matrix {
	ar := len(*a)
	ac := len((*a)[0])
	br := len(*b)
	bc := len((*b)[0])
	if ac != br {
		panic("wrong matrix type")
	}
	c := make(Matrix, ar)
	for i := 0; i < ar; i++ {
		c[i] = make([]int, bc)
	}

	ch := make(chan int)
	for i := 0; i < ar; i++ {
		go computePart(i, a, b, &c, ch)
	}
	for i := 0; i < ar; i++ {
		<-ch
	}
	return c
}

func mul(a, b *Matrix) Matrix {
	ar := len(*a)
	ac := len((*a)[0])
	br := len(*b)
	bc := len((*b)[0])
	if ac != br {
		panic("wrong matrix type")
	}
	c := make(Matrix, ar)
	for i := 0; i < ar; i++ {
		c[i] = make([]int, bc)
		for j := 0; j < bc; j++ {
			for k := 0; k < ac; k++ {
				c[i][j] += (*a)[i][k] * (*b)[k][j]
			}
		}
	}
	return c
}

// 参考
// https://qiita.com/RubyLrving/items/d4db90cbfc1a397eb139

func main() {
	N := nextInt()
	Pieces := make([]Axis, N)
	for i := 0; i < N; i++ {
		Pieces[i] = Axis{nextInt(), nextInt()}
	}
	M := nextInt()
	operations := make([]int, M)
	conversion := make([]Matrix, M+1)
	for i := 0; i <= M; i++ {
		conversion[i] = make(Matrix, 3)
		for j := 0; j < 3; j++ {
			conversion[i][j] = make([]int, 3)
		}
	}
	runtime.GOMAXPROCS(4)
	conversion[0] = Matrix{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	RotRight := Matrix{{0, 1, 0}, {-1, 0, 0}, {0, 0, 1}}
	RotLeft := Matrix{{0, -1, 0}, {1, 0, 0}, {0, 0, 1}}
	for i := 0; i < M; i++ {
		operations[i] = nextInt()
		switch operations[i] {
		case 1:
			conversion[i+1] = mul(&RotRight, &conversion[i])
		case 2:
			conversion[i+1] = mul(&RotLeft, &conversion[i])
		case 3:
			p := nextInt()
			LineX := Matrix{{-1, 0, 2 * p}, {0, 1, 0}, {0, 0, 1}}
			conversion[i+1] = mul(&LineX, &conversion[i])
		case 4:
			p := nextInt()
			LineY := Matrix{{1, 0, 0}, {0, -1, 2 * p}, {0, 0, 1}}
			conversion[i+1] = mul(&LineY, &conversion[i])
		}
	}
	Q := nextInt()
	for i := 0; i < Q; i++ {
		A, B := nextInt(), nextInt()-1
		v := Matrix{{Pieces[B].A}, {Pieces[B].B}, {1}}
		ans := mul(&conversion[A], &v)
		fmt.Println(ans[0][0], ans[1][0])
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
