package main

import (
	"bufio"
	"fmt"
	"os"
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

// 参考 ↓の③
// https://drken1215.hatenablog.com/entry/2020/04/22/002100

func main() {
	N := nextInt()
	X := nextInt() - 1
	Y := nextInt() - 1
	ans := make([]int, N)
	dp := make([][]int, N)
	for i := 0; i < N; i++ {
		dp[i] = make([]int, N)
		for j := 0; j < N; j++ {
			dp[i][j] = abs(i - j)
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			ChminInt(&dp[i][j], dp[i][X]+dp[Y][j]+1)
			ChminInt(&dp[i][j], dp[i][Y]+dp[X][j]+1)
		}
	}
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			ans[dp[i][j]]++
		}
	}
	for i := 1; i < N; i++ {
		fmt.Println(ans[i])
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

func MinOf(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}

// ChminInt 第一引数のほうが大きかった場合第二引数の値を代入する。
// 1つ目の値は参照渡しする
func ChminInt(a *int, b int) {
	if *a > b {
		*a = b
	}
}
