package main

import (
	"bufio"
	"fmt"
	"os"
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
	S, T := nextLine(), nextLine()
	ans := 0
	for i := 0; i < len(S); i++ {
		if S[i] == T[i] {
			ans++
		}
	}
	fmt.Println(ans)
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}
