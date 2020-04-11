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
	n := nextInt()
	fmt.Println(fib(n))
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

// http://www.nowhere.co.jp/blog/archives/20160909-110015.html
// ↓こっちのほうが簡潔かも
// https://atcoder.jp/contests/tenka1-2012-qualA/submissions/11755133
func fib(n int) int {
	a, b := 1, 1
	for i := 1; i < n; i++ {
		b += a
		a = b - a
	}
	return b
}
