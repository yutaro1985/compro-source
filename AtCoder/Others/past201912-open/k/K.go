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

// 行きがけ順
// https://youtu.be/1V45kF40zHc
// https://rsk0315.hatenablog.com/entry/2019/12/29/051900
// https://qiita.com/drken/items/4a7869c5e304883f539b#3-4-dfs-%E3%81%AE%E6%8E%A2%E7%B4%A2%E9%A0%86%E5%BA%8F%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6%E3%81%AE%E8%A9%B3%E7%B4%B0
// オイラーツアー
// https://www.hamayanhamayan.com/entry/2020/01/01/010045

var orgsort []int
var Parts []part
var index int

func main() {
	N := nextInt()
	org := make([][]int, N)
	Parts = make([]part, N)
	for i := 0; i < N; i++ {
		Parts[i] = part{-1, -1}
	}
	seen := make([]int, N)
	var president int
	for i := 0; i < N; i++ {
		p := nextInt() - 1
		if p < 0 {
			president = i
		} else {
			org[p] = append(org[p], i)
		}
	}
	dfs(president, org, seen)
	Q := nextInt()
	for i := 0; i < Q; i++ {
		a, b := nextInt()-1, nextInt()-1
		if b == president {
			fmt.Println("Yes")
			continue
		}
		if Parts[a].L >= Parts[b].L && Parts[a].L <= Parts[b].R {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func dfs(v int, org [][]int, seen []int) {
	if seen[v] == 0 {
		Parts[v].L = index
	}
	seen[v]++
	// fmt.Println(seen)
	orgsort = append(orgsort, v)
	for i := 0; i < len(org[v]); i++ {
		if seen[org[v][i]] == 1 {
			continue
		}
		index++
		dfs(org[v][i], org, seen)
		if i == len(org[v])-1 {
			Parts[v].R = index
		}
	}
	// for _, nextv := range org[v] {
	// 	if seen[nextv] == 1 {
	// 		continue
	// 	}
	// 	index++
	// 	dfs(nextv, index, org, seen)
	// }
}

type part struct {
	L, R int
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
