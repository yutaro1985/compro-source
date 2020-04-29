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

// golangの正規表現置換だと間に合わない
// 累積和を使う

func main() {
	N, Q := nextInt(), nextInt()
	S := nextLine()
	T := make([]int, N)
	for i := 0; i < N-1; i++ {
		var cnt int
		if S[i:i+2] == "AC" {
			cnt = 1
		}
		T[i+1] = T[i] + cnt
	}
	for i := 0; i < Q; i++ {
		l, r := nextInt(), nextInt()
		fmt.Println(T[r-1] - T[l-1])
	}
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
