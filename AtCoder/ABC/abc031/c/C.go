package main

import (
	"bufio"
	"fmt"
	"os"
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
)

func main() {
	N := nextInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = nextInt()
	}
	T := -1 << 60
	for i := 0; i < N; i++ {
		Tmax, Amax := -1<<60, -1<<60
		for j := 0; j < N; j++ {
			if i == j {
				continue
			}
			var Atmp, Ttmp int
			if i < j {
				Ta := a[i : j+1]
				for k := 0; k < len(Ta); k++ {
					switch k % 2 {
					case 0:
						Ttmp += Ta[k]
					case 1:
						Atmp += Ta[k]
					}
				}
				if Atmp > Amax {
					Amax = Atmp
					Tmax = Ttmp
				}
			} else {
				Ta := a[j : i+1]
				for k := 0; k < len(Ta); k++ {
					switch k % 2 {
					case 0:
						Ttmp += Ta[k]
					case 1:
						Atmp += Ta[k]
					}
				}
				if Atmp > Amax {
					Amax = Atmp
					Tmax = Ttmp
				}
			}
		}
		if Tmax > T {
			T = Tmax
		}
	}
	fmt.Println(T)
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

func ceil(a, b int) int {
	return (a + (b - 1)) / b
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
