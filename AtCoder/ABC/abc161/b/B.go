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
	M := nextInt()
	An := make([]int, N)
	var sum, count int
	for i := 0; i < N; i++ {
		An[i] = nextInt()
		sum += An[i]
	}
	for _, Ai := range An {
		if sum > Ai*(4*M) {
			count++
		}
	}
	if count > N-M {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
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
