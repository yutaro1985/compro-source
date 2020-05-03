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
	var n int
	// 一応、値決め打ちでなくてちゃんと計算させた場合
	for i := 0; i*i*i*i*i-(i-1)*(i-1)*(i-1)*(i-1)*(i-1) < 1e9; i++ {
		n = i + 1
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= i; j++ {
			if i*i*i*i*i-j*j*j*j*j == X {
				fmt.Println(i, j)
				return
			} else if i*i*i*i*i+j*j*j*j*j == X {
				fmt.Println(i, -j)
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
