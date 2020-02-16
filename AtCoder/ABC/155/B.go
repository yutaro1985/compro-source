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
	var N, An int
	N = nextInt()
	for i := 0; i < N; i++ {
		An = nextInt()
		if An%2 == 1 {
			continue
		}
		if An%3 != 0 && An%5 != 0 {
			fmt.Println("DENIED")
			return
		}
	}
	fmt.Println("APPROVED")
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
