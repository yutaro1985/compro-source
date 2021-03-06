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

var buf []byte = make([]byte, maxBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

func main() {
	N, Q := nextInt(), nextInt()
	ai := make([]int, N)
	for i := 0; i < Q; i++ {
		L, R, T := nextInt(), nextInt(), nextInt()
		for j := L - 1; j < R; j++ {
			ai[j] = T
		}
	}
	for _, a := range ai {
		fmt.Println(a)
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
