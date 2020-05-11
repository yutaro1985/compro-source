package main

import (
	"bufio"
	"fmt"
	"math"
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
	var n, ans, min int
	fmt.Scan(&n)
	ans = -math.MaxInt64
	min = -ans
	for i := 0; i < n; i++ {
		var R int
		// fmt.Scanのループは遅い
		sc.Scan()
		R, _ = strconv.Atoi(sc.Text())
		if ans < R-min {
			ans = R - min
		}
		if min > R {
			min = R
		}
	}
	fmt.Println(ans)
}
