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
	ans := 0
	for i := 0; i < N; i++ {
		S := nextLine()
		if i == N-1 && (S == "TAKAHASHIKUN." || S == "Takahashikun." || S == "takahashikun.") {
			ans++
		} else if S == "TAKAHASHIKUN" || S == "Takahashikun" || S == "takahashikun" {
			ans++
		}
	}
	fmt.Println(ans)
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}
