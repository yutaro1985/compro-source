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
	S := nextLine()
	numW := make([]int, N+1)
	numE := make([]int, N+1)
	for i, Si := range S {
		if Si == 'W' {
			numW[i+1] = numW[i] + 1
			numE[i+1] = numE[i]
		} else {
			numW[i+1] = numW[i]
			numE[i+1] = numE[i] + 1
		}
	}
	ans := N
	for i := 0; i < N; i++ {
		peoples := numW[i] - numW[0] + numE[N] - numE[i+1]
		if ans > peoples {
			ans = peoples
		}
	}
	fmt.Println(ans, S, numE, numW)
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
