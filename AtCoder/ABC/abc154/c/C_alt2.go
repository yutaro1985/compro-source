package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// メモリ量が増えるような書き方でトライ
// メモリ量が増えた分計算時間が減った
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
	list := make(map[int]int)
	for i := 0; i < N; i++ {
		A := nextInt()
		if _, exist := list[A]; exist {
			fmt.Println("NO")
			return
		} else {
			list[A] = A
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
