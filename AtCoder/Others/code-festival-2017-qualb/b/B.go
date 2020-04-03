package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	initialBufSize = 1e4
	maxBufSize     = 1e8 + 7
)

var buf []byte = make([]byte, maxBufSize)
var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(buf, maxBufSize)
}

func main() {
	N := nextInt()
	Di := make(map[int]int)
	for i := 0; i < N; i++ {
		Di[nextInt()]++
	}
	M := nextInt()
	for i := 0; i < M; i++ {
		D := nextInt()
		Di[D]--
		if Di[D] < 0 {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
