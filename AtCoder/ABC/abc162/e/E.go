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
	N := nextInt()
	K := nextInt()
	ans := 0
	// すべての数字が同じ場合
	ans = N * (N + 1) / 2
	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			ans += gcdof2numbers(i)
		}
		ans += gcdof2numbers(i, gcdof2numbers(b, c))
	}
	ans %= 1e9 + 7
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

func gcdof2numbers(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcdof2numbers(b, a%b)
}
