package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

const (
	initialBufSize = 10000
	maxBufSize     = 1000000009
)

var buf []byte = make([]byte, maxBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

func main() {
	N := nextInt()
	M := nextInt()
	m := make(map[int]int)
	keys := make([]int, 0, M)
	for i := 0; i < M; i++ {
		L := nextInt()
		m[L] = nextInt()
	}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	// fmt.Println(keys, m)
	maxleft := 1
	minright := N
	L := 0
	R := 0
	for _, k := range keys {
		L = k
		R = m[L]
		// 全てのL〜Rに共通する範囲を出す
		fmt.Println(L, R, maxleft, minright)
		if maxleft > R || minright < L {
			fmt.Println(0)
			return
		}
		if L < maxleft {
			L = maxleft
		} else {
			maxleft = L
		}
		if minright < R {
			R = minright
		} else {
			minright = R
		}
	}
	fmt.Println(minright, maxleft)
	fmt.Println(minright - maxleft + 1)
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

// 引数のうち最小のもの
func min(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton min() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Min(float64(res), float64(nums[i])))
	}
	return res
}

// 引数のうち最大のもの
func max(nums ...int) int {
	if len(nums) == 0 {
		panic("funciton max() requires at least one argument.")
	}
	res := nums[0]
	for i := 0; i < len(nums); i++ {
		res = int(math.Max(float64(res), float64(nums[i])))
	}
	return res
}
