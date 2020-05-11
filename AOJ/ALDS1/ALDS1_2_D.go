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

var buf []byte = make([]byte, maxBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

func main() {
	n := nextInt()
	shellSort(n)
}

func insertionSort(A []int, n, g int) ([]int, int) {
	cnt := 0
	for i := g; i < n; i++ {
		v := A[i]
		j := i - g
		for ; j >= 0 && A[j] > v; j -= g {
			A[j+g] = A[j]
			cnt++
		}
		A[j+g] = v
	}
	return A, cnt
}

func shellSort(n int) {
	cnt := 0
	var m int
	A := make([]int, n)
	for i := range A {
		A[i] = nextInt()
	}
	//適度に値に差がつくようなGを定義する
	// シェルソートの間隔の決め方
	// https://programming-place.net/ppp/contents/algorithm/sort/005.html
	G := make([]int, 0)
	for i := 1; i <= n; i = 3*i + 1 {
		G = append(G, i)
	}
	m = len(G)
	for i := len(G) - 1; i >= 0; i-- {
		var tmp int
		A, tmp = insertionSort(A, n, G[i])
		cnt += tmp
	}
	fmt.Println(m)
	sort.Sort(sort.Reverse(sort.IntSlice(G)))
	fmt.Println(strings.Trim(fmt.Sprint(G), "[]"))
	fmt.Println(cnt)
	for _, v := range A {
		fmt.Println(v)
	}
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
