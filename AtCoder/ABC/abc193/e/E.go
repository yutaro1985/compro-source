package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
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
	T := nextInt()
	for ; T > 0; T-- {
		X, Y, P, Q := nextInt(), nextInt(), nextInt(), nextInt()
		C := X + Y
		// 2C*n+X <= t <= 2C*n+X+Y
		D := P + Q
		// D*m+P <= t <= D*m+P+Q
		t := INF
		for a := X; a < X+Y; a++ {
			for b := P; b < P+Q; b++ {
				n, lcm := Crt([]int64{int64(a), int64(b)}, []int64{int64(C) * 2, int64(D)})
				if lcm == 0 {
					continue
				}
				ChminInt(&t, int(n))
			}
		}
		if t == INF {
			fmt.Println("infinity")
		} else {
			fmt.Println(t)
		}
	}
}

// Crt は中国剰余定理の実装
// common_mathを一緒に使用すること。
// https://github.com/monkukui/ac-library-go/blob/master/math/math.go
func Crt(r, m []int64) (int64, int64) {
	if len(r) != len(m) {
		panic("")
	}
	n := len(r)
	r0 := int64(0)
	m0 := int64(1)
	for i := 0; i < n; i++ {
		if !(1 <= m[i]) {
			panic("")
		}
		r1 := SafeMod(r[i], m[i])
		m1 := m[i]
		if m0 < m1 {
			r0, r1 = r1, r0
			m0, m1 = m1, m0
		}
		if m0%m1 == 0 {
			if r0%m1 != r1 {
				return 0, 0
			}
			continue
		}
		g, im := InvGcd(m0, m1)

		u1 := m1 / g
		if (r1-r0)%g != 0 {
			return 0, 0
		}

		x := (r1 - r0) / g % u1 * im % u1

		r0 += x * m0
		m0 *= u1
		if r0 < 0 {
			r0 += m0
		}
	}
	return r0, m0
}

// https://github.com/monkukui/ac-library-go/blob/master/internal/math/math.go

// @param m `1 <= m`
// @return x mod m
func SafeMod(x, m int64) int64 {
	x %= m
	if x < 0 {
		x += m
	}
	return x
}

// Fast moduler by barrett reduction
type Barrett struct {
	M  uint
	Im uint64
}

// @param m `1 <= m`
func NewBarrett(m uint) *Barrett {
	return &Barrett{
		M: m,
		// im: uint64(-1)/m + 1,
		Im: math.MaxUint64/uint64(m) + 1,
	}
}

// @return m
func (barrett *Barrett) Umod() uint {
	return barrett.M
}

// @param a `0 <= a < m`
// @param b `0 <= b < m`
// @return `a * b mod m`
func (barrett *Barrett) Mul(a, b uint) uint {
	z := uint64(a)
	z *= uint64(b)
	x, _ := bits.Mul64(z, barrett.Im)
	v := uint(z - x*uint64(barrett.M))
	if barrett.M <= v {
		v += barrett.M
	}
	return v
}

// @param n `0 <= n`
// @param m `1 <= m`
// @return `(x ** n) % m`
func PowMod(x, n int64, m int) int64 {
	if m == 1 {
		return 0
	}
	um := uint(m)
	r := uint64(1)
	y := uint64(SafeMod(x, int64(m)))

	for n > 0 {
		if n&1 > 0 {
			r = (r * y) % uint64(um)
		}
		y = (y * y) % uint64(um)
		n >>= 1
	}
	return int64(r)
}

// @param n `0 <= n`
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 || n == 7 || n == 61 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	d := int64(n - 1)
	for d%2 == 0 {
		d /= 2
	}
	for _, a := range []int64{2, 7, 61} {
		t := d
		y := PowMod(a, t, n)
		for t != int64(n-1) && y != 1 && y != int64(n-1) {
			y = y * y % int64(n)
			t <<= 1
		}
		if y != int64(n-1) && t%2 == 0 {
			return false
		}
	}
	return true
}

// @param b `1 <= b`
// @return (g, x) s.t. g = gcd(a, b), xa = g (mod b), 0 <= x < b/g
func InvGcd(a, b int64) (int64, int64) {
	a = SafeMod(a, b)
	if a == 0 {
		return b, 0
	}

	s := b
	t := a
	m0 := int64(0)
	m1 := int64(1)

	for t > 0 {
		u := s / t
		s -= t * u
		m0 -= m1 * u

		tmp := s
		s = t
		t = tmp
		tmp = m0
		m0 = m1
		m1 = tmp
	}

	if m0 < 0 {
		m0 += b / s
	}
	return s, m0
}

// @param m must be prime
// @return primitive root (and minimum in now)
func PrimitiveRoot(m int) int {
	if m == 2 {
		return 1
	}
	if m == 167772161 {
		return 3
	}
	if m == 469762049 {
		return 3
	}
	if m == 754974721 {
		return 11
	}
	if m == 998244353 {
		return 3
	}
	var divs [20]int
	divs[0] = 2
	cnt := 1
	x := (m - 1) / 2
	for x%2 == 0 {
		x /= 2
	}
	for i := 3; int64(i)*int64(i) <= int64(x); i += 2 {
		if x%i == 0 {
			divs[cnt] = i
			cnt++
			for x%i == 0 {
				x /= i
			}
		}
	}
	if x > 1 {
		divs[cnt] = x
		cnt++
	}
	for g := 2; ; g++ {
		ok := true
		for i := 0; i < cnt; i++ {
			if PowMod(int64(g), int64((m-1)/divs[i]), m) == 1 {
				ok = false
				break
			}
		}
		if ok {
			return g
		}
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
// func PowMod(a, b, m int) int {
// 	a = a % m
// 	p := 1 % m
// 	for b > 0 {
// 		if b&1 != 0 {
// 			p = (p * a) % m
// 		}
// 		b >>= 1
// 		a = (a * a) % m
// 	}
// 	return p
// }

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
