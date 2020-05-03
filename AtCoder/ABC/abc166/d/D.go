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
	X := nextInt()
	for i := -200; i < 200; i++ {
		for j := -200; j < 200; j++ {
			A5 := i * i * i * i * i
			B5 := j * j * j * j * j
			if A5-B5 == X {
				fmt.Println(i, j)
				return
			}
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
