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
	M := nextInt()
	An := make([]int, N)
	Scores := make([]int, M+1)
	for i := range An {
		An[i] = nextInt()
		Scores[An[i]]++
	}
	for i, score := range Scores {
		if score > N/2 {
			fmt.Println(i)
			return
		}
	}
	fmt.Println("?")
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
