package main

import (
	"bufio"
	"fmt"
	"os"
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
	S := nextLine()
	if S[2] == S[3] && S[4] == S[5] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}
