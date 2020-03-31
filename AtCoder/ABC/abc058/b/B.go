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
	O := nextLine()
	E := nextLine()
	ans := make([]rune, len(O)+len(E))
	for i := 0; i < len(O)+len(E); i++ {
		switch i % 2 {
		case 0:
			ans[i] = rune(O[i/2])
		case 1:
			ans[i] = rune(E[i/2])
		}
	}
	fmt.Println(string(ans))
}

func nextLine() string {
	sc.Scan()
	return sc.Text()
}
