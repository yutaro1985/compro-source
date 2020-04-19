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
	ans := 0
	for i := 0; i <= N-K+1; i++ {
		a := (K + i) * (K + i - 1) / 2
		l := (K + i) * (N*2 - K - i + 1) / 2
		nums := l - a + 1
		ans += nums
	}
	fmt.Println(ans % (1e9 + 7))
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
