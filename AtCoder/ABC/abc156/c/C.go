package main

import (
	"bufio"
	"fmt"
	"math"
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

// この解法はこれ
// https://yamakasa.net/atcoder-abc-156-c/
// TODO 公式解説PDFの方法で実装する

func main() {
	N := nextInt()
	var sum, sc, st int
	Xn := make([]int, N)
	for i := 0; i < N; i++ {
		Xn[i] = nextInt()
		sum += Xn[i]
	}
	ave := float64(sum) / float64(N)
	for i := 0; i < N; i++ {
		sc += int(math.Pow((float64(Xn[i]) - math.Ceil(ave)), 2.0))
		st += int(math.Pow((float64(Xn[i]) - math.Trunc(ave)), 2.0))
	}
	ans := MinOf(sc, st)
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
