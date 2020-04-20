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
	// いもす法を使った解
	N, K := nextInt(), nextInt()
	var ans int
	an := make([]int, N)
	csum := make([]int, N+1)
	for i := range an {
		an[i] = nextInt()
	}
	for i := 0; i < N; i++ {
		csum[i+1] = csum[i] + an[i]
	}
	for i := K; i <= N; i++ {
		ans += csum[i] - csum[i-K]
	}
	fmt.Println(ans)
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
