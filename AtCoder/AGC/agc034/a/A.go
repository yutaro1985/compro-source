package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	_, A, B, C, D := nextInt(), nextInt(), nextInt(), nextInt(), nextInt()
	S := nextLine()
	if strings.Contains(S[A-1:C], "##") || strings.Contains(S[B-1:D], "##") {
		fmt.Println("No")
		return
	}
	if C > D && !strings.Contains(S[B-2:D+1], "...") {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
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

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
