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
	sum := N * (N + 1) / 2
	n3 := N / 3
	sum3 := 3 * n3 * (n3 + 1) / 2
	n5 := N / 5
	sum5 := 5 * n5 * (n5 + 1) / 2
	n15 := N / 15
	sum15 := 15 * n15 * (n15 + 1) / 2
	fmt.Println(sum - (sum3 + sum5 - sum15))
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
