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
	N, K := nextInt(), nextInt()
	var ans int
	an := make([]int, N)
	for i := 0; i < K; i++ {
		an[i] = nextInt()
		ans += an[i]
	}
	prevsum := ans
	for i := K; i < N; i++ {
		an[i] = nextInt()
		prevsum += an[i] - an[i-K]
		ans += prevsum
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
