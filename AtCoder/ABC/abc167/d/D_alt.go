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

// 解説放送を参考に正しい実装にした
// https://youtu.be/ENSOy8u9K9I?t=2553

func main() {
	N, K := nextInt(), nextInt()
	An := make([]int, N)
	sortedAn := make([]int, 0)
	Revisit := make([]int, N+1)
	for i := range Revisit {
		Revisit[i] = -1
	}
	for i := range An {
		An[i] = nextInt()
	}
	cur := 1
	for Revisit[cur] == -1 {
		Revisit[cur] = len(sortedAn)
		sortedAn = append(sortedAn, cur)
		cur = An[cur-1]
	}
	looplength := len(sortedAn) - Revisit[cur]
	if K <= Revisit[cur] {
		fmt.Println(sortedAn[K])
	} else {
		K -= Revisit[cur]
		K %= looplength
		fmt.Println(sortedAn[K+Revisit[cur]])
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
