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

var buf []byte = make([]byte, maxBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

func main() {
	N := nextInt()
	T := nextInt()
	Dn := make([]int, N)
	var sumA, sumB int
	for i := 0; i < N; i++ {
		Ai, Bi := nextInt(), nextInt()
		sumA += Ai
		sumB += Bi
		Dn[i] = Bi - Ai
	}
	if sumB > T {
		fmt.Println(-1)
		return
	}
	sort.Ints(Dn)

	// Nを含めないとWAになる
	for i := 0; i <= N; i++ {
		if sumA <= T {
			fmt.Println(i)
			return
		}
		sumA += Dn[i]
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
