package main

import (
	"bufio"
	"fmt"
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
	K := nextInt()
	An := make([]int, N)
	Cn := make([]int, N*(N-1)/2)
	index := 0
	for i := 0; i < N; i++ {
		An[i] = nextInt()
	}
	for i := 1; i < N; i++ {
		for j := i; j < N; j++ {
			Cn[index] = An[i-1] * An[j]
			index++
		}
	}
	sort.Ints(Cn)
	fmt.Println(Cn[K-1])
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
