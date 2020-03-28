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

func main() {
	K := nextInt()
	N := nextInt()
	An := make([]int, N)
	for i, _ := range An {
		An[i] = nextInt()
	}
	ans := K - (An[0] + (K - An[N-1]))
	for i := 1; i < N-1; i++ {
		ans = MinOf(ans, K-(An[i]-An[i-1]), K-(An[i+1]-An[i]))
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

func MinOf(vars ...int) int {
	min := vars[0]
	for _, i := range vars {
		if min > i {
			min = i
		}
	}
	return min
}
