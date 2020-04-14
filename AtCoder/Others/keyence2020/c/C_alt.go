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
	S := nextInt()
	// そのまま表示。こちらのほうが遅い。
	for i := 0; i < N; i++ {
		if i < K {
			pf(S, i, N)
		} else if S != 1e9 {
			pf(S+1, i, N)
		} else {
			pf(1, i, N)
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

func pf(s int, i int, N int) {
	if i != N {
		fmt.Printf("%d ", s)
	} else {
		fmt.Printf("%d\n", s)
	}
}
