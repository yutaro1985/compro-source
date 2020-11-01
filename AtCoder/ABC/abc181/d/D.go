package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	initialBufSize = 1e4
	maxBufSize     = 1e9 + 7
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
	S := nextLine()
	if len(S) == 1 {
		if S == "8" {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
		return
	}
	if len(S) == 2 {
		Snum, _ := strconv.Atoi(S)
		Snumre, _ := strconv.Atoi(string(S[1]) + string(S[0]))
		if Snum%8 == 0 || Snumre%8 == 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
		return
	}
	list := make(map[string]bool)
	numscnt := make(map[string]int)
	num := 112
	for num < 1000 {
		nums := strconv.Itoa(num)
		if !strings.Contains(nums, "0") {
			list[nums] = true
		}
		num += 8
	}
	for i := 0; i < len(S); i++ {
		numscnt[string(S[i])]++
	}
	for k := range list {
		check := make(map[string]int)
		for i := 0; i < 3; i++ {
			check[string(k[i])]++
		}
		isMul8 := true
		for num, ckcnt := range check {
			if _, e := numscnt[num]; !e {
				isMul8 = false
				break
			}
			if ckcnt > numscnt[num] {
				isMul8 = false
				break
			}
		}
		if isMul8 {
			fmt.Println("Yes")
			return
		}
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

// ChminInt 第一引数のほうが大きかった場合第二引数の値を代入する。
// 1つ目の値は参照渡しする
func ChminInt(a *int, b int) {
	if *a > b {
		*a = b
	}
}

// ChmanInt 第一引数のほうが小さかった場合第二引数の値を代入する。
// 1つ目の値は参照渡しする
func ChmaxInt(a *int, b int) {
	if *a < b {
		*a = b
	}
}
