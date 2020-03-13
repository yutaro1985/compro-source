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
	N := nextInt()
	K := nextInt()
	Q := nextInt()
	correctanswers := make([]int, N)
	for i := 0; i < Q; i++ {
		Ai := nextInt() - 1
		correctanswers[Ai]++
	}
	for i := 0; i < N; i++ {
		score := K - Q + correctanswers[i]
		if score > 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
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
