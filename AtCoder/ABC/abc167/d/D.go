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
	An := make([]int, N)
	visitPlace := make([]int, N+1)
	firstPlace := map[int]int{}
	repeatPlace := 0
	for i := 0; i < N; i++ {
		An[i] = nextInt()
	}
	visitPlace[0] = 1
	for i := 1; i < N+1; i++ {
		visitPlace[i] = An[visitPlace[i-1]-1]
		if _, exist := firstPlace[visitPlace[i]]; exist {
			repeatPlace = i
			break
		} else {
			firstPlace[visitPlace[i]] = i
		}
	}
	first := firstPlace[visitPlace[repeatPlace]]
	if K <= first {
		fmt.Println(visitPlace[K])
	} else {
		K -= first - 1
		K %= repeatPlace - first
		repeatTowns := visitPlace[first:repeatPlace]
		index := K - 1
		if index < 0 {
			index = len(repeatTowns) - 1
		}
		fmt.Println(repeatTowns[index])
	}
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
